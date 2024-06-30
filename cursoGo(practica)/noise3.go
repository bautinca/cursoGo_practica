// Aca vamos a seguir viendo ruido pero ahora vamos a usar HILOS! recordar que los hilos comparten la memoria de
// un proceso. Ahora bien, nosotros vamos a ver hilos especiales en GO llamadas GORRUTINAS

// Pero no son hilos comunes, sino que tienen algunas diferencias:

// HILOS COMUNES
//	1. Las maneja el SO
//	2. La creacion de un hilo si bien es menos costoso que crear un proceso nuevo, sigue siendo costoso en terminos de recursos
//	3. Cada hilo tiene su propio stack de memoria
//	4. Cada hilo es una entidad en el SO y el scheduler del SO decide cuanto tiempo de CPU le asigna
//	5. El scheduler del SO puede mover los hilos entre distintas CPUs segun le convenga
//	6. La sincronizacion entre hilos se suele llevar a cabo con locks y semaforos
//	7. Sincronizarlos puede causar errores como deadlocks o condiciones de carrera

// GORRUTINAS
//	1. Son puramente manejadas por GO, no por el SO
//	2. Las gorrutinas tambien se les llama HILOS LIGEROS, porque justamente existen dentro de un mismo hilo
//	3. La creacion de las Gorrutinas son muy baratas y rapidas
//	4. El stack de las gorrutinas es dinamico y no algo estatico reservado como los hilos comunes
//	5. Tenemos varias gorrutinas dentro de un mismo hilo
//	6. Un solo hilo entonces puede ejecutar multiples gorrutinas
//	7. Para sincronizar gorrutinas tenemos el concepto de CANALES proporcionado por Go
//	8. Los canales facilitan la sincronizacion entre gorrutinas

// Bueno entonces usaremos las GORRUTINAS para trabajar concurrentemente con el ruido, antes que nada traeremos todo lo que hicimos
// en noise2.go

package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type pos struct {
	x, y float32
}

const winWidth, winHeight int = 800, 600

type color struct {
	r, g, b byte
}

func setPixel(x, y int, c color, pixels []byte) {

	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

func lerp(b1, b2 byte, pct float32) byte {
	return byte(float32(b1) + pct*(float32(b2)-float32(b1)))
}

func colorLerp(c1, c2 color, pct float32) color {
	return color{lerp(c1.r, c2.r, pct), lerp(c1.g, c2.g, pct), lerp(c1.b, c2.b, pct)}
}

func getGradient(c1, c2 color) []color {
	result := make([]color, 256)

	for i, _ := range result {
		pct := float32(i) / float32(255)
		result[i] = colorLerp(c1, c2, pct)
	}
	return result
}

func getDualGradient(c1, c2, c3, c4 color) []color {
	result := make([]color, 256)

	for i, _ := range result {
		pct := float32(i) / float32(255)

		if pct < 0.5 {
			result[i] = colorLerp(c1, c2, pct*float32(2))
		} else {
			result[i] = colorLerp(c3, c4, pct*float32(1.5)-float32(0.5))
		}
	}
	return result
}

func clamp(min, max, value int) int {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}

func rescaleAndDraw(noise []float32, min, max float32, gradient []color, pixels []byte) {

	scale := 255.0 / (max - min)
	offset := min * scale

	for i := range noise {
		noise[i] = noise[i]*scale - offset

		c := gradient[clamp(0, 255, int(noise[i]))]
		p := i * 4
		pixels[p] = c.r
		pixels[p+1] = c.g
		pixels[p+2] = c.b
	}

}

func turbulence(x, y, frequency, lacunarity, gain float32, octaves int) float32 {
	var sum float32
	amplitude := float32(1.0)
	for i := 0; i < octaves; i++ {
		f := snoise2(x*frequency, y*frequency) * amplitude
		if f < 0 {
			f = -1.0 * f
		}
		sum += f
		frequency = frequency * lacunarity
		amplitude = amplitude * gain
	}
	return sum
}

func fbm2(x, y, frequency, lacunarity, gain float32, octaves int) float32 {

	var sum float32
	amplitude := float32(1.0)
	for i := 0; i < octaves; i++ {
		sum += snoise2(x*frequency, y*frequency) * amplitude
		frequency = frequency * lacunarity
		amplitude = amplitude * gain
	}

	return sum
}

// Ahora esta funcion toma el ancho y alto de la ventana ya que nos servira para tomar las coordenadas (x,y)
func makeNoise2(pixels []byte, frequency, lacunarity, gain float32, octaves, w, h int) {
	// Bueno aca es donde creamos el ruido en toda la ventana entonces lo primero que tenemos que ver es decidir
	// cuantas gorrutinas diferentes queremos ejecutar cuando generamos ruido

	// Resulta que al momento de generar ruido como lo hacemos en esta funcion makeNoise2() es un problema concurrente
	// Lo bueno es que a este ruido es facil de aplicarle concurrencia

	// Tambien se podria paralelizar si tenemos varias CPUs por lo que se ejecutan simultaneamente varias instrucciones
	// entonces lo primero que haremos es consultar cuantas CPUs tenemos, esto siempre es util antes de aplicar cualquier gorrutina
	numRoutines := runtime.NumCPU()

	// Aca generamos wg que es un grupo de espera para justamente esperar a q se terminen de ejecutar todas las gorrutinas
	// para proseguir con la ejecucion del hilo principal
	// La declaramos...
	var wg sync.WaitGroup

	// Ahora le decimos cuantas gorrutinas esperara, es logico que esperara numRoutines gorrutinas
	wg.Add(numRoutines)

	// Tambien lo que agregaremos es el famoso MUTEX=LOCK!!, esto lo veremos luego donde lo usaremos pero recordemos
	// que sirve para que una gorrutina tome el lock y ejecute un bloque de codigo en forma atomica, es decir, el lock garantiza atomicidad para un bloque de codigo
	// y asi evitar condiciones de carrera
	var mutex = &sync.Mutex{}

	min := float32(9999.0)
	max := float32(-9999.0)
	noise := make([]float32, winWidth*winHeight)

	// Ahora bien, nosotros tenemos el arreglo 'noise' que es justamente un arreglo de valores ruido de la ventana 'pixels'
	// Entonces antes que nada lo que haremos es dividir el arreglo 'noise' en numRoutines partes iguales para que justamente
	// Cada CPU trabaje una porcion del arreglo 'noise' de manera justa
	batchSize := len(noise) / numRoutines

	// Luego haremos un for para que en cada iteracion crear una nueva gorrutina
	for i := 0; i < numRoutines; i++ {
		// Para crear la gorrutina usamos la palabra clave 'go', y ademas usaremos funciones LAMBDA, que recordar que son funciones anonimas
		// por ejemplo el map() en python!
		go func(i int) {
			// Entonces apenas generamos una gorrutina colocaremos un defer para que justamente al terminan de ejecutarse
			// esta gorrutina entonces marcar una tilde en la variable 'wg', que es justamente la que espera que finalicen gorrutinas
			defer wg.Done() // Entonces al terminar de ejecutarse esta gorrutina que generariamos marcamos como una gorrutina completada en la variable 'wg'

			// Lo que haremos es tomar el inicio en el arreglo noise para cada CPU, es decir si por ejemplo
			// el arreglo noise es de 250 elementos entonces una CPU tendria el inicio en indice 0, y la otra en el indice 125
			start := i * batchSize // Este seria el indice inicio para cada CPU
			// Asi como definimos el inicio tambien tenemos que definir el fin del arreglo noise para cada CPU
			end := start + batchSize - 1

			// Ya teniendo el inicio y fin para cada CPU ahora si hacemos el loop desde el inicio hasta el fin de ese pedazo de arreglo
			for j := start; j < end; j++ {
				// Ahora definimos las coordenadas x e y
				x := j % w // Justamente para no escaparnos del ancho de la ventana la coordenada X
				y := (j - x) / h
				// Ahora si generaremos el valor ruido en esa coordenada (x,y) para ello tomaremos el Turbulence
				noise[j] = turbulence(float32(x), float32(y), frequency, lacunarity, gain, octaves)

				// Y ahora redefinimos el minimo y maximo de los valores que nos da turbulence() para obtener el valor minimo y maximo
				// osea tal cual lo que teniamos antes para escalarlo luego
				// Y ojo, ACA USAREMOS EL MUTEX! ¿Porque? Justamente para prevenir una condicion de carrera
				mutex.Lock() // A partir de aca una gorrutina puede tomar el lock
				if noise[j] < min {
					min = noise[j]
				} else if noise[j] > max {
					max = noise[j]
				}
				mutex.Unlock() // Aca la gorrutina suelta el lock y la pone en disponibilidad para otra gorrutina
				// Ahora bien, ojo con los LOCKS porque si abusamos de ellos entonces el uso de gorrutinas no tendria sentido, ya que justamente
				// lo que garantizan las gorrutinas es atomicidad y una ejecucion SECUENCIAL por parte de un hilo ligero
				// Por lo tanto, si tenemos muchisimos locks entonces es directamente lo mismo que no tener gorrutinas ya que ejecutamos todo secuencialmente
				// Asi que ojo con el uso de los locks o mutex, es decir, porque el exceso de locks termina arruinando el recurso
				// de concurrencia
			}
		}(i) // Colocamos esta 'i' donde justamente representa el parametro que le pasamos a la gorrutina, es decir
		// se generarian distintas gorrutinas donde a c/u le paso el valor de i cuando hacemos la iteracion
	}
	// Si ejecutamos las gorrutinas asi como estan entonces al ver la ventana vamos a ver algo raro, como que no se termina de ejecutar todo el ruido por toda la ventana
	// sino que hay rangos de la ventana donde no se genera ruido, y esto es logico que pase porque los distintas gorrutinas no estan coordinadas, es decir, mientras las gorrutinas se ejecutan el hilo principal sigue
	// su ejecucion. Es decir, para corregir esto lo ideal seria que el for de las numRoutines se termine de ejecutar todo, es decir, antes de seguir con el hilo principal
	// que las distintas gorrutinas terminen su ejecucion. Entonces al no poner esta restriccion es logico que al momento de ejecutar el programa
	// no se termine de generar el ruido en todas las coordenadas (x,y)

	// Entonces, tenemos que hacer algo para que todas las gorrutinas terminen de ejecutarse antes de proseguir con el hilo principal y ejecutar las proximas
	// 2 lineas getDualGradient() y rescaleAndDraw(). Para ello despues de definir el numero de CPUs lo que haremos es declarar una variable
	// wg de grupo de espera. Entonces lo que le pasaremos es la cantidad de cosas que esperara a que terminen de ejecutarse para proseguir con el hilo principal
	// por tanto es logico que a wg le pasaremos todas las gorrutinas que generamos segun la cantidad de CPUs que tenemos

	// Entonces aca al hilo principal le marcamos que espere la ejecucion de las gorrutinas para proseguir
	wg.Wait()

	// Entonces ahora al ejecutar el programa si se va a poder visuaizar bien todo el ruido en la ventana ya que terminan de ejecutarse todas
	// las gorrutinas antes de proseguir con el hilo principal

	gradient := getDualGradient(color{0, 0, 175}, color{80, 160, 244}, color{12, 192, 75}, color{255, 255, 255})
	rescaleAndDraw(noise, min, max, gradient, pixels)
}

// -------------------------------------------------------------------

func main() {

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	frequency := float32(0.01)
	gain := float32(0.2)
	lacunarity := float32(3.0)
	octaves := 3

	makeNoise2(pixels, frequency, lacunarity, gain, octaves, winWidth, winHeight)

	keyState := sdl.GetKeyboardState()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {

			case *sdl.QuitEvent:
				return
			}
		}

		mult := 1

		if keyState[sdl.SCANCODE_LSHIFT] != 0 || keyState[sdl.SCANCODE_RSHIFT] != 0 {
			mult = -1
		}

		if keyState[sdl.SCANCODE_O] != 0 {
			octaves = octaves + 1*mult
			makeNoise2(pixels, frequency, lacunarity, gain, octaves, winWidth, winHeight)
		}

		if keyState[sdl.SCANCODE_F] != 0 {
			frequency = frequency + 0.001*float32(mult)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves, winWidth, winHeight)
		}

		if keyState[sdl.SCANCODE_G] != 0 {
			gain = gain + 0.1*float32(mult)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves, winWidth, winHeight)
		}

		if keyState[sdl.SCANCODE_A] != 0 {
			lacunarity = lacunarity + 0.1*float32(lacunarity)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves, winWidth, winHeight)
		}

		pixelsPointer := unsafe.Pointer(&pixels[0])
		tex.Update(nil, pixelsPointer, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16)

	}
}

// ------------------------------SIMPLEX NOISE----------------------------------------------

// Aca pegaremos la implementacion en codigo ya hecha del Simplex Noise, lo que si no la estudiaremos a detalle porque es muy de bajo nivel
// y no nos interesa, lo que si nos va a interesar sera como usar la funcion Simplex Noise en este caso 2D, pero si queremos saber mas estas funciones
// las podemos implementar nosotros mismos, pero para ir a lo practico ya importamos una elaborada

/* This code ported to Go from Stefan Gustavson's C implementation, his comments follow:
 * https://github.com/stegu/perlin-noise/blob/master/src/simplexnoise1234.c
 * SimplexNoise1234, Simplex noise with true analytic
 * derivative in 1D to 4D.
 *
 * Author: Stefan Gustavson, 2003-2005
 * Contact: stefan.gustavson@liu.se
 *
 *
 * This code was GPL licensed until February 2011.
 * As the original author of this code, I hereby
 * release it into the public domain.
 * Please feel free to use it for whatever you want.
 * Credit is appreciated where appropriate, and I also
 * appreciate being told where this code finds any use,
 * but you may do as you like.
 */

/*
 * This implementation is "Simplex Noise" as presented by
 * Ken Perlin at a relatively obscure and not often cited course
 * session "Real-Time Shading" at Siggraph 2001 (before real
 * time shading actually took off), under the title "hardware noise".
 * The 3D function is numerically equivalent to his Java reference
 * code available in the PDF course notes, although I re-implemented
 * it from scratch to get more readable code. The 1D, 2D and 4D cases
 * were implemented from scratch by me from Ken Perlin's text.
 *
 * This file has no dependencies on any other file, not even its own
 * header file. The header file is made for use by external code only.
 */

func fastFloor(x float32) int {
	if float32(int(x)) <= x {
		return int(x)
	}
	return int(x) - 1
}

// Static data

/*
 * Permutation table. This is just a random jumble of all numbers 0-255
 * This needs to be exactly the same for all instances on all platforms,
 * so it's easiest to just keep it as static explicit data.
 * This also removes the need for any initialisation of this class.
 *
 */

// Esta seria la tabla de PERMUTACION, son numeros del 0 al 255 organizados aleatoriamente
var perm = [256]uint8{151, 160, 137, 91, 90, 15,
	131, 13, 201, 95, 96, 53, 194, 233, 7, 225, 140, 36, 103, 30, 69, 142, 8, 99, 37, 240, 21, 10, 23,
	190, 6, 148, 247, 120, 234, 75, 0, 26, 197, 62, 94, 252, 219, 203, 117, 35, 11, 32, 57, 177, 33,
	88, 237, 149, 56, 87, 174, 20, 125, 136, 171, 168, 68, 175, 74, 165, 71, 134, 139, 48, 27, 166,
	77, 146, 158, 231, 83, 111, 229, 122, 60, 211, 133, 230, 220, 105, 92, 41, 55, 46, 245, 40, 244,
	102, 143, 54, 65, 25, 63, 161, 1, 216, 80, 73, 209, 76, 132, 187, 208, 89, 18, 169, 200, 196,
	135, 130, 116, 188, 159, 86, 164, 100, 109, 198, 173, 186, 3, 64, 52, 217, 226, 250, 124, 123,
	5, 202, 38, 147, 118, 126, 255, 82, 85, 212, 207, 206, 59, 227, 47, 16, 58, 17, 182, 189, 28, 42,
	223, 183, 170, 213, 119, 248, 152, 2, 44, 154, 163, 70, 221, 153, 101, 155, 167, 43, 172, 9,
	129, 22, 39, 253, 19, 98, 108, 110, 79, 113, 224, 232, 178, 185, 112, 104, 218, 246, 97, 228,
	251, 34, 242, 193, 238, 210, 144, 12, 191, 179, 162, 241, 81, 51, 145, 235, 249, 14, 239, 107,
	49, 192, 214, 31, 181, 199, 106, 157, 184, 84, 204, 176, 115, 121, 50, 45, 127, 4, 150, 254,
	138, 236, 205, 93, 222, 114, 67, 29, 24, 72, 243, 141, 128, 195, 78, 66, 215, 61, 156, 180}

//---------------------------------------------------------------------

// Funcion de Gradiente
func grad2(hash uint8, x, y float32) float32 {
	h := hash & 7 // Convert low 3 bits of hash code
	u := y
	v := 2 * x
	if h < 4 {
		u = x
		v = 2 * y
	} // into 8 simple gradient directions,
	// and compute the dot product with (x,y).

	if h&1 != 0 {
		u = -u
	}
	if h&2 != 0 {
		v = -v
	}
	return u + v
}

// Funcion Simplex Noise en 2D, vemos que le damos coordenadas (X,Y) y nos devuelve un float32, claramente
// como la Simplex Noise es 2D le tenemos q dar 2 parametros, osea X e Y y nos devuelven el punto aleatorio Noise, osea
// un valor de ruido nos retorna
func snoise2(x, y float32) float32 {

	const F2 float32 = 0.366025403 // F2 = 0.5*(sqrt(3.0)-1.0)
	const G2 float32 = 0.211324865 // G2 = (3.0-Math.sqrt(3.0))/6.0

	var n0, n1, n2 float32 // Noise contributions from the three corners

	// Skew the input space to determine which simplex cell we're in
	s := (x + y) * F2 // Hairy factor for 2D
	xs := x + s
	ys := y + s
	i := fastFloor(xs)
	j := fastFloor(ys)

	t := float32(i+j) * G2
	X0 := float32(i) - t // Unskew the cell origin back to (x,y) space
	Y0 := float32(j) - t
	x0 := x - X0 // The x,y distances from the cell origin
	y0 := y - Y0

	// For the 2D case, the simplex shape is an equilateral triangle.
	// Determine which simplex we are in.
	var i1, j1 uint8 // Offsets for second (middle) corner of simplex in (i,j) coords
	if x0 > y0 {
		i1 = 1
		j1 = 0
	} else { // lower triangle, XY order: (0,0)->(1,0)->(1,1)
		i1 = 0
		j1 = 1
	} // upper triangle, YX order: (0,0)->(0,1)->(1,1)

	// A step of (1,0) in (i,j) means a step of (1-c,-c) in (x,y), and
	// a step of (0,1) in (i,j) means a step of (-c,1-c) in (x,y), where
	// c = (3-sqrt(3))/6

	x1 := x0 - float32(i1) + G2 // Offsets for middle corner in (x,y) unskewed coords
	y1 := y0 - float32(j1) + G2
	x2 := x0 - 1.0 + 2.0*G2 // Offsets for last corner in (x,y) unskewed coords
	y2 := y0 - 1.0 + 2.0*G2

	// Wrap the integer indices at 256, to avoid indexing perm[] out of bounds
	ii := uint8(i)
	jj := uint8(j)

	// Calculate the contribution from the three corners
	t0 := 0.5 - x0*x0 - y0*y0
	if t0 < 0.0 {
		n0 = 0.0
	} else {
		t0 *= t0
		n0 = t0 * t0 * grad2(perm[ii+perm[jj]], x0, y0)
	}

	t1 := 0.5 - x1*x1 - y1*y1
	if t1 < 0.0 {
		n1 = 0.0
	} else {
		t1 *= t1
		n1 = t1 * t1 * grad2(perm[ii+i1+perm[jj+j1]], x1, y1)
	}

	t2 := 0.5 - x2*x2 - y2*y2
	if t2 < 0.0 {
		n2 = 0.0
	} else {
		t2 *= t2
		n2 = t2 * t2 * grad2(perm[ii+1+perm[jj+1]], x2, y2)
	}

	// Add contributions from each corner to get the final noise value.
	return (n0 + n1 + n2)
}
