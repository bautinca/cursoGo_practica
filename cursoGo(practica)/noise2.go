package main

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

// Hasta ahora ya vimos Simplex Noise y Fractal Noise y como jugar con sus parametros, pero y el color que onda? Porque hasta ahora hicimos todo ruido en blanco y negro
// El primer paso es hacer una interpolacion lineal entre 2 colores para ello haremos una funcion llamada colorLerp() donde recibe 2 colores y el porcentaje

// Recordar que Interpolacion lineal hicimos en pong2.go para justamente ubicar el puntaje de ambos jugadores, recordar que la interpolacion
// lineal siempre toma 2 valores y un valor porcentual que seria la interpolacion entre estos 2 valores
// Para ello haremos una simple funcion lerp()
// Esto nos dara la interpolacion lineal entre 2 bytes
func lerp(b1, b2 byte, pct float32) byte {
	return byte(float32(b1) + pct*(float32(b2)-float32(b1)))
}

// Pero nosotros queremos hacer la interpolacion lineal entre 2 colores, pero recordar q cada color esta representado por 4 parametros
// donde cada parametro es 1 byte, por tanto la interpolacion entre 2 colores es equivalente a la interpolacion entre los 4 parametros de cada color RGBA
// Logicamente el Lerp entre colores me da otro color que este color si seria justamente el de color ya que hacemos una interpolacion entre 2 colores blanco y negro
func colorLerp(c1, c2 color, pct float32) color {
	return color{lerp(c1.r, c2.r, pct), lerp(c1.g, c2.g, pct), lerp(c1.b, c2.b, pct)}
}

// Ahora hacemos una funcion para obtener el gradiente entre 2 colores y devolvemos un arreglo de colores, es decir, si por ejemplo c1=rojo y c2=azul entonces
// esta funcion devolveria un arreglo de 256 colores donde los colores van pasando del rojo al azul a modo degradado
func getGradient(c1, c2 color) []color {
	// Ahora lo que haremos sera un arreglo de colores de capacidad 256, donde justamente pasamos del color rojo al color azul a modo degradado en este arreglo de colores
	result := make([]color, 256)

	// Ahora lo que hacemos es iterar por el arreglo para ir llenandolo de colores entre el rojo y el azul
	for i, _ := range result {
		// A cada valor de arreglo lo dividimos por 255 y diremos que es el valor porcentual, hacemos esto para justamente pasarlo a porcentaje
		pct := float32(i) / float32(255)
		// Luego llenamos el color en dicha posicion del arreglo, para ello usaremos colorLerp() que justamente le pasamos un extremo del color(por ejemplo rojo) y el otro extremo del color (por ejemplo azul)
		// Y con colorLerp() nos devuelve el intermedio entre dichos colores pero teniendo en cuenta el defasaje que proporciona el valor porcentual que lo definimos antes
		result[i] = colorLerp(c1, c2, pct)
	}
	// Finalmente retornamos el arreglo de degradados de rojo a azul por ejemplo
	return result
}

// Aca haremos una funcion como getGradient() pero esta recibe solo 2 colores para obtener el gradiente entre dichos colores, pero lo cierto es que puedo
// obtener el gradiente entre varios colores por ejemplo 4 colores
func getDualGradient(c1, c2, c3, c4 color) []color {
	result := make([]color, 256)

	for i, _ := range result {
		pct := float32(i) / float32(255)

		// Aca cambiaria, si el valor porcentual es menor a 0.5 entonces haremos el gradiente entre los 2 primeros colores c1 y c2
		// pero ojo, el valor porcentual lo multiplicamos por 2 para justamente obtener todo el rango de colores entre los colores c1 y c2
		if pct < 0.5 {
			result[i] = colorLerp(c1, c2, pct*float32(2))
			// En cambio si el valor porcentual es >= 0.5 haremos el gradiente entre los otros 2 colores c3 y c4
		} else {
			result[i] = colorLerp(c3, c4, pct*float32(1.5)-float32(0.5))
		}
	}
	return result
}

// Ahora definimos una funcion clamp() que lo que hace es tomar un minimo y maximo y algun valor. Lo que hace la funcion
// clamp() es limitar un valor a un rango especifico, asegura que el valor no decaiga por debajo del minimo ni por encima del maximo
func clamp(min, max, value int) int {
	if value < min {
		value = min
	} else if value > max {
		value = max
	}
	return value
}

// Ahora la funcion rescaleAndDraw() ademas recibe un gradiente que es un arreglo de colores
func rescaleAndDraw(noise []float32, min, max float32, gradient []color, pixels []byte) {

	scale := 255.0 / (max - min)
	offset := min * scale

	for i := range noise {
		noise[i] = noise[i]*scale - offset
		// Cuando escalamos el valor ruido lo que haremos es del valor escalado es primero hacer que el valor este entre 0 y 255 con clamp()
		// luego lo que haremos es que dicho resultado sea para tomar un color del arreglo de colores gradient
		// Luego cuando ya tomamos el color dicho color lo usaremos para pintar justamente los pixeles de la ventana pixels
		c := gradient[clamp(0, 255, int(noise[i]))] // Aca tomamos un color
		p := i * 4
		pixels[p] = c.r // Y coloreamos los pixeles con ese color que tomamos del arreglo de colores gradient
		pixels[p+1] = c.g
		pixels[p+2] = c.b
	}

}

// Haremos otra tecnica llamada TURBULENCIA, donde recibe tal cual los mismos parametros que el Fractal Noise
// Es decir, es otro tipo de ruido donde logicamente devuelve un valor ruido float32
// Este ruido tambien se hace con ayuda del Simplex Noise
func turbulence(x, y, frequency, lacunarity, gain float32, octaves int) float32 {
	// Al principio lo haremos igual al Fractla Noise es decir tenemos una variable sum que es lo q retornaremos
	var sum float32
	// Luego una amplitud la definimos como fija, igual que el fractal noise 1.0
	amplitude := float32(1.0)
	// Ahora iteramos por las octavas
	for i := 0; i < octaves; i++ {
		// Llamamos a Simplex Noise para que nos de su valor de ruido
		f := snoise2(x*frequency, y*frequency) * amplitude
		// Ahora si el resultado 'f' es menor a 0 entonces lo redefinimos como con su valor absoluto
		if f < 0 {
			f = -1.0 * f
		}
		// Finalmente el valor ruido 'f' lo sumamos a 'sum'
		sum += f
		// La frecuencia la redefinimos multiplicandola por el lacunarity
		frequency = frequency * lacunarity
		// Y redefinimos la amplitud
		amplitude = amplitude * gain
	}
	// Finalmente devolvemos 'sum' que tiene el valor de ruido para las coordenadas (x,y)
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

func makeNoise2(pixels []byte, frequency, lacunarity, gain float32, octaves int) {

	noise := make([]float32, winWidth*winHeight)

	i := 0

	min := float32(9999.0)
	max := float32(-9999.0)

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {

			// noise[i] = fbm2(float32(x), float32(y), frequency, lacunarity, gain, octaves)

			// Aca ahora en lugar de que la coordenada (x,y) su ruido sea producido por el Fractal Noise fbm2() haremos
			// que el ruido sea producido por la turbulencia
			noise[i] = turbulence(float32(x), float32(y), frequency, lacunarity, gain, octaves)

			if noise[i] < min {
				min = noise[i]
			} else if noise[i] > max {
				max = noise[i]
			}

			i++
		}
	}
	// Aca definimos el gradiente, lo definiremos entre rojo y azul, entonces logicamente
	// getGradient() nos devuelve un arreglo de colores de gradiente, justamente todos los colores en el arreglo van del rojo al azul
	// a modo degradado, eso es justamente lo que hace getGradient() todo gracias a la funcion colorLerp() que devuelve el color intermedio entre 2 colores
	// con ayuda de un valor porcentual
	//gradient := getGradient(color{255, 0, 0}, color{0, 0, 255})
	//rescaleAndDraw(noise, min, max, gradient, pixels)

	// Aca usaremos el getDualGradient para ahora utilzar el gradiente entre 4 colores!
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

	makeNoise2(pixels, frequency, lacunarity, gain, octaves)

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
			makeNoise2(pixels, frequency, lacunarity, gain, octaves)
		}

		if keyState[sdl.SCANCODE_F] != 0 {
			frequency = frequency + 0.001*float32(mult)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves)
		}

		if keyState[sdl.SCANCODE_G] != 0 {
			gain = gain + 0.1*float32(mult)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves)
		}

		if keyState[sdl.SCANCODE_A] != 0 {
			lacunarity = lacunarity + 0.1*float32(lacunarity)
			makeNoise2(pixels, frequency, lacunarity, gain, octaves)
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
