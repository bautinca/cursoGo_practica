// Vamos a usar la biblioteca SDL2 (Simple DirectMedia Layer 2), es una biblioteca multiplataforma
// diseñada para proporcionar una capa de abstraccion de hardware de bajo nivel sobre graficos, sonido y entrada. Es muy
// popular en el desarrollo de videojuegos y apps multimedia porque simplifica muchas tareas comunes al interactuar con el hardware del sist.

package main

// Ahora vamos a importar el sig. paquete para ver cosas basicas de SDL2
import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// Antes que nada vamos a definir constantes que definirian el tamaño de la ventana en PIXELES
const winWidth, winHeight int = 800, 600

// Este es el struct RGBA que representaria el color de cada pixel
type color struct {
	r, g, b byte
}

// No leer esta funcion que dibuja un pixel hasta que en el main() se lo indique:
// Ahora si vamos a crear una funcion simple que dibuja un pixel, logicamente este pixel
// recibira su posicion en x, en y y ademas su COLOR!, que logicamente su color estara definido por los parametros
// RGBA donde cada parametro habiamos mencionado que pesa 1 byte, pero este RGBA sera un struct
// Luego por ultimo indicaremos que reciba la matriz de pixeles que seria la ventana
func setPixel(x, y int, c color, pixels []byte) {
	// Vamos a setear el indice, osea donde pintaremos, entonces supongamos que tenemos la sig, matriz:
	//		____________
	//	    |_0_|_1_|_2_|
	//      |_3_|_4_|_5_|
	//		|_6_|_7_|_8_|

	// Y queremos ir hacia el centro, osea 4, por tanto X=1 y Y=1, entonces si Y=1
	// entonces tendriamos que ir a la 2° fila, solor para ir a la segunda fila y posicionarnos en el 3 tendriamos que hacer
	// (y*winWidth), pero como queremos ir al 4 entonces le sumamos la coordenada X, es decir (y*winWidth + x)
	// Entonces ahi ya nos posicionamos en el 4
	index := (y*winWidth + x) * 4
	// Ahora seleccionamos ese casillero en la matrix 'pixels' y le seteamos el color
	// recordar que para setear el color tenemos el struct 'c' donde el primer componente es R(Red), el sig. G(Green),
	// luego B(Blue) y A(Alpha)

	// Vamos a asegurarnos que el index este dentro del rango de los pixeles, osea:
	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r   // Siguiendo el ejemplo anterior el 4 seria el componente rojo (R)
		pixels[index+1] = c.g // un pixel despues, el 5 seria el componente verde (G)
		pixels[index+2] = c.b // finalmente un pixel siguiente, el 6 seria el componente blue (B)
		// Por ahora ignoramos alpha
	}

}

func main() {

	// Haciendo uso de SDL2 vamos a crear una simple ventana, vamos a ver la primer funcion que es
	// CreateWindow() que es justamnte para crear una ventana donde el primer parametro es el titulo, y luego
	// los sig. 2 parametros es la posicion X e Y de la ventana, luego el ancho y alto de la ventana en pixeles y luego
	// el ultimo parametro son flags, vamos a colocar que se muestre la ventana simplemente

	// Otra cosa que no aclaramos es que CreateWindow() devuelve 2 valores de retorno, el segundo valor de retorno es por si sucede algun error
	// en la creacion de la ventana
	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)

	// Chequeamos si tenemos algun error en la creacion de la ventana
	if err != nil {
		fmt.Println(err)
		return
	}

	// Ahora logicamente cuando abrimos algo que tenemos que hacer despues? CERRARLO! Es por ello que si abrimos una ventana
	// la tenemos que cerrar, esto lo hacemos con Destroy() para justamente liberar los recursos. Pero para asegurarnos que esta funcion se ejecute
	// apenas antes de llegar al return aplicamos la instruccion reservada de GO que es DEFER, recordemos que defer servia para ejecutar
	// cosas apenas antes de llegar al return de la funcion que contiene al defer, osea main()
	defer window.Destroy()

	// Ahora ya creamos la ventana 'window', lo que haremos sera dibujar cosas en ella!, para ello usaremos la funcion
	// CreateRenderer() que funciona igual que CreateWindow() y tambien devuelve un segundo valor que es por si hubo algun error
	// El primer parametro que recibe es el mas importante y es la ventana a la que se asociara, que es logicamente 'window'
	// Luego el segundo parametro por el momento colocaremos -1 y luego algunos flags, el unico que usaremos es RENDER_ACCELERATED
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	// Chequeamos obviamente si sucedio algun error
	if err != nil {
		fmt.Println(err)
		return
	}

	// Luego al igual que la ventana, el renderizador tambien se debe destruir
	defer renderer.Destroy()

	// Ya tenemos el renderizador, ahora vamos a hacer por fin una textura dentro de la ventana
	// logicamente para crear una textura sigue el mismo patron que para crear una ventana y un render, osea CreateTexture()
	// El primer parametro es un formato de pixel ¿Como formato de pixel? Si formato de pixel
	// Uno puede pensar a los pixeles como una serie de bytes donde sabemos que el 255 es blanco y el 0 es negro
	// por ejemplo [8,8,2,255,0, etc] es decir, esa secuencia de bytes serian los pixeles, el pixel cuyo color es 8, etc.
	// Y lo que le debemos decir a SDL2 es que diseño de pixeles queremos, nosotros vamos a usar el formato RGBA (RED - GREEN - BLUE - ALPHA)
	// El parametro 'Alpha' regularia la transparencia del pixel, entonces por ejemplo si tenemos:

	//		[RGBA] = [255, 0, 0, 0] = ROJO PURO ya que tenemos 255 que seria el maximo en R(Red)
	//		[RGBA] = [255, 255, 255, 0] = BLANCO

	// Entonces nosotros vamos a usar el formato ABGR8888 justamente el '8' es porque cada numero representaria 1 byte en el arreglo
	// osea si tengo [255, 0, 0, 0] el 255 es un byte, el 0 otro byte, el 0 otro byte, y asi...

	// Luego el segundo parametro simplemente colocamos TEXTUREACCESS_STREAMING

	// Luego los otros 2 parametros serian el ancho y alto de la ventana donde se crearia la textura, osea winWidth y winHeight
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))

	// Verificamos si sucede error en la creacion de la textura
	if err != nil {
		fmt.Println(err)
		return
	}

	// El destructor obviamente de la textura con un defer para indicar que obviamente siempre se ejecute antes de llegar al return del main
	defer tex.Destroy()

	// Bien ya creamos una ventana, un renderizador y una textura. Ahora la gran pregunta es ¿Como dibujamos cosas en la ventana? Necesitamos pixeles
	// Para entender bien como dibujar imaginemos que tenemos un buffer de 0 a 8 es decir:

	// [___|___|___|___|___|___|___|___|___]
	//   0   1   2   3   4   5   6   7   8

	// Ahora bien, imaginemos que tenemos una computadora muuuy antigua donde su resolucion es de 3x3 pixeles por tanto:
	//		____________
	//	    |_0_|_1_|_2_|
	//      |_3_|_4_|_5_|
	//		|_6_|_7_|_8_|

	// Ahora imaginemos que en el buffer tenemos colocares donde cada color lo representamos por sus iniciales por ingles
	// para que se entienda bien, ejemplo R(Red), B(Black), etc...

	// [_R_|_B_|_Y_|_W_|_etc_|___|___|___|___]
	//   0   1   2   3    4    5   6   7   8

	// Claramente cada color en el buffer ocupa 1 byte por asi quisimos
	// Entonces en la matriz 3x3 los colores se representarian tal cual asi igual que en el buffer, osea:
	//		____________
	//	    |_R_|_B_|_Y_|
	//      |_W_|etc|_5_|
	//		|_6_|_7_|_8_|

	// Entonces lo que nosotros haremos es crear una MATRIZ DE PIXELES donde lo que hacemos es contar la cantidad total de pixeles
	// que tenemos en la ventana, si la ventana es de 800*600 es porque justamente tiene esa cantidad de pixeles ancho*alto

	// Pero el ancho*alto tambien lo multiplicamos por 4 ¿Porque? Porque estamos en el formato ABGR!, osea RGBA = [red, green, blue, alpha] donde
	// dijimos que c/u ocuparia 1 byte, por lo tanto cada pixel ocuparia 4 bytes ya que tiene esos parametros, osea el grado de rojo que tienen, el grado de verde que tienen
	// el grado de azul, y el grado de alpha
	// Por lo tanto la matriz de pixeles quedaria:
	pixels := make([]byte, winWidth*winHeight*4)

	// Ahora si leer la funcion para crear un pixel que definimos antes del main()

	// Una vez leida la funcion setPixel() que es para pintar algunos pixeles ahora la vamos a usar
	// vamos a recorrer toda la matriz de pixeles
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{255, 0, 0}, pixels) // Basicamente lo queqestamos haciendo es recorrer por todos los pixeles de la ventana y en todos los pixeles
			// estamos estableciendo el color rojo ya que [RGB] = [255, 0, 0]
		}
	}

	// Ahora lo que falta es asociar esto a nuestra TEXTURA (tex) entonces haremos
	// tenemos que usar la funcion Update() donde el primer parametro por ahoral o dejamos en nil
	// pero el segundo le pasamos la matriz y el tercero es el ancho de la ventana multiplicado por la cantidad de bytes de cada pixel

	// Pero antes que nada a pixels debemos hacerle lo sig.
	pixelsPointer := unsafe.Pointer(&pixels[0])

	// Ahora si la funcion Update() que mencionamos antes
	tex.Update(nil, pixelsPointer, winWidth*4)

	// Ahora tenemos que usar nuestro RENDERIZADOR usando la funcion COPY() pasadnole como primer parametro
	// la textura (tex) y los otros 2 parametros por ahora en nil
	renderer.Copy(tex, nil, nil)

	// Y finalmente presentamos la renderizacion en la ventana
	renderer.Present()

	// Luego colocamos un delay para que la ventana no se cierre inmediatamente ya que sin este delay entonces apenas ejecuta el programa
	// se crea la ventana y un microsegundo despues se cierra ya que la funcion main llega a su return
	// Asi que colocaremos aqui un delay de 2 segundos (el parametro que recibe son milisegundos)
	sdl.Delay(2000)

	return
}
