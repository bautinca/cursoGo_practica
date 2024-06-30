package main

// Vamos a hablar de NOISE (Ruido) ¿Que son las funciones Noise? Lo podemos imaginar como un grafico cartesiano en 2D con los ejes Y y X
// Donde el ruido hace zigzag en direccion del eje X, es como que se desarrolla un mamarracho a lo largo de la recta X, bueno eso es RUIDO

// Va a haber ocasiones en que sea necesario generar ruido de una manera fluida. Para tener este ejecto de un ruido
// fluido usaremos algo llamado COHERENTE. Ahora vamos a decir 2 propiedades:

//	1. Si una entrada esta cerca de otra, la salida tambien lo estara
//	2. SI las entradas estan muy alejadas entonces las salidas seran aleatorias

// Por ejemplo si hacemos un grafico sinusoidal sen(x) vemos que es regular, es decir, las entradas
// que estan mas alejadas son predecibles, siempre obtenemos lo mismo

// Entonces como podemos hacer Noise Coherente? Dicho sea de paso el 'ruido coherente' en programacion
// se refiere a un tipo de ruido que tiene propiedades especificas que lo hacen util en la generacion de texturas, paisajes, graficos
// y otros efectos visuales en computacion grafica. A diferencia del Noise aleatorio puro, que carece de estructura y continuidad, el ruido coherente
// posee cierta suavidad y continuidad, lo que lo hace mas natural y esteticamente agradable

// El VALUE NOISE (Ruido valor) es un tipo de ruido coherente. Se caracteriza por asignar valores aleatorios a puntos de una rejilla y luego interpolar estos valores de manera suave para obtener una apariencia
// continua y natural, a diferencia de otros tipos de ruido. Este tipo de ruido no utiliza gradientes como otros, sino directamente en valores
// asignados a los puntos de la rejilla

// En el Value Noise imaginemos que tenemos una recta en X, y en dicha recta encima dibujamos valores de X a intervalos regulares
// Osea:

//				-----|-----|-----|-----|-----|-----|-----|----> X

// Luego en vada valor de X colocamos un valor aleatorio en Y

//											 .
//					 .	                           .
//				-----|-----|-----|-----|-----|-----|----> X
//					       .     .

// Luego para llenar los espacios entre los puntos hacemos la INTERPOLACION LINEAL donde lo que hace
// es basicamente unir los puntos por una RECTA, donde no lo dibujaremos porque es dificil de dibujarlo pero se entiende, basicamente
// unir los puntos con rectas. Bueno acabamos de hacer un tipo de ruido coherente que se llama VALUE NOISE

// Luego hay otro tipo de ruido coherente que es el GRADIENT NOISE, donde inicialmente es lo mismo que el VALUE NOISE, osea en el eje X elegimos intervalos
// regulares:

//				-----|-----|-----|-----|-----|-----|-----|----> X

// Pero a diferencia del VALUE NOISE, el GRADIENT NOISE en lugar de elegir un valor aleatorio en eje Y lo que hace es elegir una PENDIENTE aleatoria en cada
// valor en el eje X. Entonces en lugar de hacer una interpolacion lineal entre 2 puntos lo que hace es interpolar pero entre pendientes por lo que se obtienen entre pendientes
// curvas suabes que respetan justamente la inclinacion de todas las pendientes

// Cabe aclarar que para los ejemplos en el GRADIENT NOISE y en el VALUE NOISE Estamos viendo todo UNIDIMENSIONAL, pero esta logica se extiende para N dimensiones

// Ahora vamos a ver otro tipo de ruido que es el mas famoso de todos que es el PERLIN NOISE, es un tipo de RUIDO COHERENTE desarrollado en 1983.
// Tambien hay otro ruido coherente que se llama SIMPLEX NOISE que es con el que nosotros trabajaremos sobre el codigo de pong2.go por tanto lo copiamos y lo pegamos en este archivo

// La funcion de simplex noise ya estan hechas en varios lugares, por ejemplo nos podemos encontrar la funcion de Simplex Noise en Github pero para codigo en C, lo que haremos sera copiar la implementacion en codigo
// del Simplex Noise y pegarla aca abajo de todo

// -------------------------------------------------------------------------

import (
	"fmt"
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

// -----------------------------------------------------------------

// Aca vamos a hacer la funcion de reescalado por cada valor que nos devuelve la funcion snoise2() y coloree el pixel
// Es decir, recibe logicamente la matriz de valores ruido 'noise', el valor 'min' y 'max' que representaria el valor maximo de valor ruido de la matriz noise
// y el valor minimo de valor ruido de la matriz noise, y obviamente la ventana 'pixels' para poder dibujar en ella
func rescaleAndDraw(noise []float32, min, max float32, pixels []byte) {
	// Nosotros lo que vamos a querer hacer es escalar todos los valores ruido de la matriz 'noise' entre valores
	// 0 y 255, sabiendo que el maximo valor de la matriz noise es 'max' y el minimo valor de la matriz noise es 'min'
	scale := 255.0 / (max - min)
	offset := min * scale

	// Ahora iteramos por todos los valores de la matriz
	for i := range noise {
		noise[i] = noise[i]*scale - offset // Aca aplicamos el escalado
		b := byte(noise[i])                // El escalado lo pasamos a tipo de dato byte
		pixels[i*4] = b                    // Luego el valor ya escalado entre 0 y 255 y pasado a tipo de dato byte lo podemos finalmente
		// colocar en la ventana 'pixels', pero para ello recordar que lo tenemos que multiplicar por 4 ya que
		// cada pixel tiene 4 parametros RGBA (Red, Green, Blue, Alpha)
		pixels[i*4+1] = b
		pixels[i*4+2] = b
	}

	// Y listo! ya tenemos la funcion que reescala los valores ruido entre 0 y 255 y los pasa a tipo de dato
	// bytes en la matriz de pixeles pixels

}

// -----------------------------FRACTAL NOISE--------------------------------
// Lo que haremos sera implementar otro ruido, sera el FRACTAL NOISE, nosotros con lo que trabajabamos
// con snoise2() es el SIMPLEX NOISE, pero ahora implementaremos el FRACTAL
// La funcion esta tomara las coordenadas (x,y) del pixel obviamente, toma una frecuencia, lacunarity, gain y octaves
// Todos estos parametros son propios del FRACTAL NOISE, son como reguladores del tipo de ruido
// Logicamente la funcion devuelve un valor ruido
func fbm2(x, y, frequency, lacunarity, gain float32, octaves int) float32 {
	// La implementacion del Fractal Noise se ve asi
	var sum float32
	// Tambien tendremos una amplitud que empieza a partir del 1
	amplitude := float32(1.0)
	// Luego lo que haremos es iterar por c/u de las octavas y aplicar snoise2() que el valor ruido que me devuelve
	// se lo voy sumando a 'sum', pero las coordenadas (x,y) la multiplicamos por la frecuencia y el valor ruido le multiplicamos la amplitud
	for i := 0; i < octaves; i++ {
		sum += snoise2(x*frequency, y*frequency) * amplitude
		// Tambien por cada iteracion la frecuencia va cambiando donde en cada iteracion le vamos multiplicando por el lacunarity
		frequency = frequency * lacunarity
		// Y en cada iteracion tambien va cambiando la amplitud donde la vamos multiplicando por gain
		amplitude = amplitude * gain
	}
	// Finalmente cuando todo este hecho retornamos 'sum'
	return sum
}

// -----------------------------------------------------------------------

// Lo que haremos aca sera implementar una funcion para usar Simplex Noise fullscreen para ver que onda
// como funciona
func makeNoise(pixels []byte) {
	// Lo primero que haremos sera crear una matriz de float32 que tenga el tamaño de la ventana
	// Y donde justamente en esta matriz se almacenara los valores ruido que nos devuelva la funcion snoise2
	noise := make([]float32, winWidth*winHeight)

	// Declaramos y definimos el indice de la matriz
	i := 0

	// El valor minimo y maximo para poder escalar los valores ruido
	min := float32(9999.0)
	max := float32(-9999.0)
	// Lo que vamos a hacer sera iterar por todos los pixeles de la ventana asi nomas
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			// Cuando pasamos punto por punto (es decir pixel por pixel) de la totalidad de la ventana
			// Lo que haremos sera aplicarle simplex noise a cada coordenada usando la funcion importada
			// Almacenaremos el valor resultado en la matriz

			// Lo mas maravilloso de todo es que podemos jugar con la frecuencia, es decir, yo le puedo pasar
			// a snoise2() las coordenadas (x,y) del pixel, pero puedo jugar y pasarle las coordenadas pero divididas por 100 por ejemplo, y esto me va generando
			// otras texturas en la pantalla, por ejemplo coloquemos /100 ambos parametros
			//noise[i] = snoise2(float32(x)/100, float32(y)/100)

			// Aca usaremos en lugar del snoise2() el fbm2() que implementamos, osea el Fractal Noise con ayuda del Simplex Noise (snoise2)
			// La diferencia ahora con snoise2() es que el Fractal Noise me pide parametros extra y no solo las coordenadas (x,y)
			// Tengo que definir los parametros frequency, amplitude, gain y octaves
			// Decidimos que:
			//	Frequency = 0.001
			// 	Lacunarity = 2
			//	Gain = 2
			//	Octaves = 3
			noise[i] = fbm2(float32(x), float32(y), 0.001, 2, 2, 3)

			// Entonces ahi lo podemos ejecutar fbm2() y el resultado con esos parametros hardcodeados nos dara como una mancha de zebra nublado
			// El resultdo es bastante similar a que si usabamos snoise2() solamente, pero ¿Como juegan estos parametros adicionales que  tiene el Fractal Nose
			// fbm2()?

			// Para entender bien como juegan estos parametros frequency, lacunarity, gain, octaves lo que haremos sera crearnos
			// un mecanismo para que a traves del teclado poder regular estos parametros mediante el uso del teclado en tiempo de ejecucion

			// Si el valor ruido almacenado es menor al minimo entonces redefinimos el 'min' por el valor
			// ruido
			// En cambio si el valor ruido es mayor al max entonces redefinimos 'max' por el valor del ruido, osea:
			if noise[i] < min {
				min = noise[i]
			} else if noise[i] > max {
				max = noise[i]
			}
			// Para que hacemos todo esto de los IFs? Para que justamente 'min' quede con el valor minimo
			// de todos los valores ruido y 'max' quede con el valor maximo de todos los valores ruido
			// Y asi teniendo el valor ruido minimo y maximo poder escalar todo los valores ruido en ese rango

			// Sumamos +1 al indice i
			i++
		}
	}

	// Ahora si aca aplicaremos la funcion que implementamos antes rescaleAndDraw() que lo que hace es reescalar
	// los valores ruido entre 0 y 255, los pasa a tipo de dato byte y los coloca en la matriz de pixeles pixels
	rescaleAndDraw(noise, min, max, pixels)
}

// Esta funcion hace lo mismo que makeNoise() la unica diferencia es que makeNoise2() recibe los parametros del Fractal Noise
func makeNoise2(pixels []byte, frequency, lacunarity, gain float32, octaves int) {

	noise := make([]float32, winWidth*winHeight)

	i := 0

	min := float32(9999.0)
	max := float32(-9999.0)

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {

			noise[i] = fbm2(float32(x), float32(y), frequency, lacunarity, gain, octaves)

			if noise[i] < min {
				min = noise[i]
			} else if noise[i] > max {
				max = noise[i]
			}

			i++
		}
	}

	rescaleAndDraw(noise, min, max, pixels)
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
	// Aca establecemos los valores iniciales de los parametros del Fractal Noise fbm2()
	frequency := float32(0.01)
	gain := float32(0.2)
	lacunarity := float32(3.0)
	octaves := 3

	// --------------------------------------------------------
	// Y aca llamamos a makeNoise() que implementamos como funcion para que pase por todos los pixeles
	// de la ventana
	// makeNoise(pixels)

	// Al correr la funcion asi nomas vamos a ver que toda la ventan se pone en negor y como colocamos
	// en la implementacion de la funcion makeNoise() un print entonces se van a imprimir todos los retornos flotantes
	// que nos devuelve la funcion snoise2() por STDOUT, osea por la Shell

	// Ahora si vemos los valores que se imprimen logicamente se van a imprimir tantos valores por cuantos pixeles tengamos en la ventana
	// ya que justamente interamos por todos los pixeles de la ventana y por cada pixel devolvemos un valor de ruido
	// Y estos valores de ruido estan en un rango entre numeros negativos y numeros positivos pero muy proximos al 0
	// Entonces estos valores ruido que se imprimen los podemos escalar para que sean en el rango que deseemos

	// Ahora bien, como los escalaremos? Como nuestra matriz de pixeles 'pixels' es una matriz de tipo de dato byte eso quiere decir que
	// son valores entre 0 y 255, por tanto los podemos escalar en el intervalo [0, 255]

	// Si nos fijamos a makeNoise() le pasamos la matriz 'pixels' entonces lo que va a suceder en el primerisimo frame de todos es que se va a generar
	// ruido en  toda la ventana, donde ahora cada pixel tiene un valor entre 0 y 255, por lo que representaria un color

	// ¿ Que ruido generaria? Bueno como si fuera una TV sin señal
	// ----------------------------------------------------------

	// Aca implementamos nuevamente makeNoise() pero la cambiaremos ligeramente, la llamaremos makeNoise2() que es tal cual lo mismo que makeNoise() la unica
	// diferencia es que recibe los parametros del Fractal Noise frequency, gain, lacunarity y octaves
	makeNoise2(pixels, frequency, lacunarity, gain, octaves)

	// Aca implementamos un estado del teclado para justamente regular los parametros del Fractal Noise fbm2()
	keyState := sdl.GetKeyboardState()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {

			case *sdl.QuitEvent:
				return
			}
		}

		// Lo que vamos a querer hacer aca es ir cambiando todos los parametros del Fractal Noise
		mult := 1
		// Diremos que si el usuario presiona el LSHIFT o el RSHIFT entonces configuramos la variable 'mult'
		if keyState[sdl.SCANCODE_LSHIFT] != 0 || keyState[sdl.SCANCODE_RSHIFT] != 0 {
			mult = -1
		}
		// Luego elegimos otra tecla para otro parametros del fbm2() para ajustar sus parametros, por ejemplo
		// ahora elegimos la tecla 'O' para justamente regular el parametro 'octaves' de fbm2()
		if keyState[sdl.SCANCODE_O] != 0 {
			octaves = octaves + 1*mult                               // Entonces si mantenemos presionado SHIFT y apretamos 'O' se va restando mult, en cambio si solo apretamos 'O' se va sumando mult
			makeNoise2(pixels, frequency, lacunarity, gain, octaves) // Y llamamos a la funcion makeNoise2() para ver los resultados en la ventana
		}
		// Ahora para otro parametro 'F' para regular la frecuencia de fbm2()
		if keyState[sdl.SCANCODE_F] != 0 {
			frequency = frequency + 0.001*float32(mult)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves)
		}
		// Ahora para otro parametro 'G' para regular el gain de fbm2()
		if keyState[sdl.SCANCODE_G] != 0 {
			gain = gain + 0.1*float32(mult)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves)
		}
		// Ahora para el ultimo parametro 'L' de Lacunarity
		if keyState[sdl.SCANCODE_A] != 0 {
			lacunarity = lacunarity + 0.1*float32(lacunarity)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves)
		}

		// Finalmente si corremos todo vamos a ver que en la ventana por cada fotograma vamos a podere ir ajustando distintos parametros
		// del Fractal Noise haciendo uso del teclado:
		//	O = Para ajustar Octaves
		//	F = Para ajustar Frequency
		//	G = Para ajustar Gain
		//	L = Para ajustar Lacunarity

		// Sumado a que si mantenemos SHIFT mientras presionamos las letras es para reducir la intensidad de dicho parametro
		// por ejemplo SHIFT + G es para ir reduciendo el Gain por fotograma

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
