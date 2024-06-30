// Aca lo que vamos a hacer es jugar al juego pong pero de fondo un ruido, para ello primero copiamos y pegamos aca todo el contenido de pong2.go que era lo ultimo q hicimos del juego
package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// Ahora aca vamos a copiar y pegar todo lo de noise3.go asi crudamente pero cambiando algunos detalles minimos para que justamente se pueda implementar con el juego pong
//---------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------------------------------------------------------------------------
//---------------------------------------------------------------------------------------------------------------------

// Lo primero que agregaremos en todo el codigo que teniamos en noise3.go es agregar un estado de ruido
// esto lo hacemos para q el usuario pueda elegir entre elegir un Fractal Noise o un Turbulence Noise
type NoiseType int

const (
	FBM NoiseType = iota
	TURBULENCE
)

// A lerp() la renombramos como nlerp() ya que tenemos otra funcion lerp() para la implementacion del juego pong
func nlerp(b1, b2 byte, pct float32) byte {
	return byte(float32(b1) + pct*(float32(b2)-float32(b1)))
}

func colorLerp(c1, c2 color, pct float32) color {
	return color{nlerp(c1.r, c2.r, pct), nlerp(c1.g, c2.g, pct), nlerp(c1.b, c2.b, pct)}
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

// Tambien modificaremos esta funcion, no recibe el argumento pixels sino que devuelve una nueva
// ventana de pixeles modificados y ademas el ancho y alto de la ventana
func rescaleAndDraw(noise []float32, min, max float32, gradient []color, w, h int) []byte {

	// Aca colocaremos lo que retornaremos, tenemos que hacer un arreglo de ancho*alto que nos pasa el usuario
	result := make([]byte, w*h*4) // Lo multiplicamos *4 porque cada pixel esta compuesto por 4 colores RGBA

	scale := 255.0 / (max - min)
	offset := min * scale

	for i := range noise {
		noise[i] = noise[i]*scale - offset

		c := gradient[clamp(0, 255, int(noise[i]))]
		p := i * 4
		result[p] = c.r
		result[p+1] = c.g // Pintamos el pixel
		result[p+2] = c.b
	}
	// Retornamos el resultado de pintar los pixeles
	return result

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

// Lo primero que vamos a cambiar un poco es esta funcion makeNoise2()
// Una de las cosas q le agregaremos es el tipo de sonido q selecciona el usuario ya que justamente la funcion makeNoise2() es la encargada de generar ruido
// Por tanto recibe un selector de ruido como primer parametro
// Ademas otra de las modificaciones es q retorna 3 valores!, la matriz ruido, el minimo y maximo del ruido generado en dicha matriz!
func makeNoise2(noiseType NoiseType, frequency, lacunarity, gain float32, octaves, w, h int) (noise []float32, min, max float32) {

	numRoutines := runtime.NumCPU()

	var wg sync.WaitGroup

	wg.Add(numRoutines)

	var mutex = &sync.Mutex{}

	min = float32(9999.0)
	max = float32(-9999.0)
	noise = make([]float32, w*h) // Aca el tamaño del arreglo noise lo define el usuario segun que ancho y que alto quiere que tenga el ruido

	batchSize := len(noise) / numRoutines

	for i := 0; i < numRoutines; i++ {

		go func(i int) {

			defer wg.Done()

			start := i * batchSize

			end := start + batchSize - 1

			for j := start; j < end; j++ {

				x := j % w
				y := (j - x) / h

				// Aca diremos que si el usuario eligio un noiseType == FBM entonces aplicamos el ruido FBM2, en cambio si eligio
				// TURBULENCE entonces aplicamos el otro ruido
				if noiseType == TURBULENCE {
					noise[j] = turbulence(float32(x), float32(y), frequency, lacunarity, gain, octaves)
				} else if noiseType == FBM {
					noise[j] = fbm2(float32(x), float32(y), frequency, lacunarity, gain, octaves)
				}

				mutex.Lock()
				if noise[j] < min {
					min = noise[j]
				} else if noise[j] > max {
					max = noise[j]
				}
				mutex.Unlock()
			}
		}(i)
	}
	wg.Wait()
	// Tambien lo q modificaremos es que solamente retorne el arreglo de ruido
	return noise, min, max
}

//------------------------------SIMPLEX NOISE-----------------------

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

//--------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------
//--------------------------------------------------------------------------------------------------------------------

type ball struct {
	pos    pos
	radius float32
	xv     float32
	yv     float32
	color  color
}

type pos struct {
	x, y float32
}

type paddle struct {
	pos   pos
	w     float32
	h     float32
	score int
	color color
}

func lerp(a float32, b float32, pct float32) float32 {
	return a + pct*(b-a)
}

func (paddle *paddle) draw(pixels []byte) {

	startX := int(paddle.pos.x - paddle.w/2)
	startY := int(paddle.pos.y - paddle.h/2)

	for y := 0; y < int(paddle.h); y++ {
		for x := 0; x < int(paddle.w); x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}

	numX := lerp(paddle.pos.x, float32(winWidth)/2, 0.2)

	drawNumber(pos{numX, 35}, paddle.color, 10, paddle.score, pixels)

}

func (ball *ball) draw(pixels []byte) {

	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.pos.x+x), int(ball.pos.y+y), ball.color, pixels)
			}
		}
	}
}

func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {
	ball.pos.x += ball.xv
	ball.pos.y += ball.yv

	if ball.pos.y-ball.radius < 0 || ball.pos.y+ball.radius > float32(winHeight) {
		ball.yv = -ball.yv
	}

	if ball.pos.x < 0 {
		rightPaddle.score++

		ball.pos.x = float32(winWidth) / 2
		ball.pos.y = float32(winHeight) / 2

		state = start
	}

	if ball.pos.x > float32(winWidth) {
		leftPaddle.score++
		ball.pos.x = float32(winWidth) / 2
		ball.pos.y = float32(winHeight) / 2

		state = start

	}

	if ball.pos.x-float32(ball.radius) < leftPaddle.pos.x+leftPaddle.w/2 {
		if ball.pos.y > leftPaddle.pos.y-leftPaddle.h/2 && ball.pos.y < leftPaddle.pos.y+leftPaddle.h/2 {
			ball.xv = -ball.xv
			ball.pos.x = leftPaddle.pos.x + float32(leftPaddle.w)/2.0 + float32(ball.radius)
		}
	}

	if ball.pos.x+float32(ball.radius) > rightPaddle.pos.x-rightPaddle.w/2 {
		if ball.pos.y > rightPaddle.pos.y-rightPaddle.h/2 && ball.pos.y < rightPaddle.pos.y+rightPaddle.h/2 {
			ball.xv = -ball.xv
			ball.pos.x = rightPaddle.pos.x - float32(rightPaddle.w)/2.0 - float32(ball.radius)
		}
	}
}

func (paddle *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		if paddle.pos.y-float32(paddle.h)/2 >= 0 {
			paddle.pos.y -= 7
		}
	}

	if keyState[sdl.SCANCODE_DOWN] != 0 {

		if paddle.pos.y+float32(paddle.h)/2 <= float32(winHeight) {
			paddle.pos.y += 7
		}
	}
}

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.pos.y = ball.pos.y
}

func clear(pixels []byte) {

	for i := range pixels {
		pixels[i] = 0
	}
}

const winWidth, winHeight int = 800, 600

type gameState int

const (
	start gameState = iota
	play
)

var state = start

var nums = [][]byte{
	{
		1, 1, 1,
		1, 0, 1, // El numero 0
		1, 0, 1,
		1, 0, 1,
		1, 1, 1,
	},
	{
		1, 1, 0,
		0, 1, 0, // Numero 1
		0, 1, 0,
		0, 1, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1, // El numero 2
		1, 1, 1,
		1, 0, 0,
		1, 1, 1,
	},
	{
		1, 1, 1,
		0, 0, 1,
		0, 1, 1, // El numero 3
		0, 0, 1,
		1, 1, 1,
	},
}

func drawNumber(pos pos, color color, size int, num int, pixels []byte) {

	startX := int(pos.x) - (size*3)/2

	startY := int(pos.y) - (size*5)/2

	for i, v := range nums[num] {

		if v == 1 {

			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					setPixel(x, y, color, pixels)
				}
			}
		}

		startX += size

		if (i+1)%3 == 0 {
			startY += size
			startX -= size * 3
		}
	}
}

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

	player1 := paddle{pos{100, 200}, 20, 100, 0, color{255, 255, 255}}
	player2 := paddle{pos{700, 100}, 20, 100, 0, color{255, 255, 255}}

	ball := ball{pos{float32(winWidth) / 2, float32(winHeight) / 2}, 20, 7, 7, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	// Aca generamos el arreglo de ruido, el minimo y maximo ruido generado en dicho arreglo noise
	noise, min, max := makeNoise2(FBM, 0.01, 0.2, 2, 3, winWidth, winHeight)

	// Ahora teniendo el arreglo de ruido lo q haremos es aplicar el reescalado para dibujar reescaleAndDraw()
	// pero antes debemos generar el gradiente entre 2 colores, lo haremos del rojo al negro
	gradient := getGradient(color{255, 0, 0}, color{0, 0, 0})

	// Luego ahora si aplicamos el reescalado y dibujamos, pero no dibujamos en la ventana posta pixels,
	// sino que dibujamos en otro arreglo aparte que nos da como resultado
	noisePixels := rescaleAndDraw(noise, min, max, gradient, winWidth, winHeight)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {

			case *sdl.QuitEvent:
				return
			}
		}

		if state == play {

			player1.update(keyState)
			player2.aiUpdate(&ball)
			ball.update(&player1, &player2)

		} else if state == start {

			if keyState[sdl.SCANCODE_SPACE] != 0 {

				if player1.score == 3 || player2.score == 3 {
					player1.score = 0
					player2.score = 0
				}
				state = play
			}
		}

		clear(pixels)

		// Ahora lo que hacemos es en cada fotograma dibujar el ruido en la ventana, por eso ni bien se limpia la ventana
		// lo que haremos es dibujar la ventana con  el ruido para luego dibujar las barras y la pelota, y todo lo q tenga q ver con el juego
		// Dibujamos antes el fondo que el juego pong ya que con esto conseguimos que las barras, el puntaje, la pelota, y todo lo q tenga q ver con el juego
		// se dibuje por encima del dibujado del ruido

		// Pero recordar que el arreeglo noisePixels tiene los pixeles de ruido por tanto deberiamos iterar por dicho arreglo noisePixels y cambiar el valor del arreglo pixels
		// por el color de los pixeles de noisePixels para asi pintar toda la ventana
		for i, _ := range noisePixels {
			pixels[i] = noisePixels[i]
		}

		player1.draw(pixels)
		ball.draw(pixels)
		player2.draw(pixels)

		pixelsPointer := unsafe.Pointer(&pixels[0])
		tex.Update(nil, pixelsPointer, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		sdl.Delay(16)

	}
}
