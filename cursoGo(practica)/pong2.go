package main

// Vamos a seguir con el juego PONG refinandolo un poco mas, pero la base ya la tenemos

// Lo primero que vamos a hacer es ELIMINAR LOS CASTEOS porque en pong.go tenemos varios casteos en varias partes del codigo
// y a veces confunde cuando hay que castear y cuando no, es por ello que ahora vamos a declarar TODOS los datos como FLOAT32

// Luego lo que vamos a hacer es agregar el PUNTAJE!!, es decir, si alguien anota u otro
import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type ball struct {
	pos    pos
	radius float32 // Aca cambiamos por float32
	xv     float32
	yv     float32
	color  color
}

type pos struct {
	x, y float32
}

type paddle struct {
	pos pos
	w   float32 // Aca cambiamos por float32
	h   float32 // Aca cambiamos por float32
	// Agregamos un nuevo atributo que sea 'score' que es el puntaje de la paleta
	score int
	color color
}

// Ahora hacemos una funcion lerp() que tiene que ver con los puntajes que van haciendo las paletas
// esta funcion recibira 2 numeros y un porcentaje, lo que hace la funcion lerp() es calcular el valor medio entre 2 numeros
// la definicion matematica de lerp es INTERPOLACION LINEAL y su formula es lerp(a,b,t) = a + t(b-a)
// 'a' es el valor inicial, 'b' el valor final y 't' es un factor de interpolacion que siempre se encuentra en el rango [0,1], si t=0 entonces
// la fucion retorna 'a' y si t=1 la funcion retorna 'b', el valor intermedio de 't' logicamente retorna el valor intermedio entre 'a' y 'b'
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

	// Ahora vamos a dibujar el score de las paletas, ¿Donde lo dibujaremos? Bueno aca
	// haremos uso de lerp() para establecer la coordenada en X donde dibujaremos el puntaje de la paleta
	// para ello dibujaremos el score entre la coordenada X de la paleta (que es su pixel central) y el centro de la ventana pero con
	// un ligero defasaje de 0.2 mas cerca del centro de la ventana, es decir...
	numX := lerp(paddle.pos.x, float32(winWidth)/2, 0.2)
	// Ahora teniendo la coordenada en X del centro del score que dibujaremos ahora si dibujamos el score
	drawNumber(pos{numX, 35}, paddle.color, 10, paddle.score, pixels)
	// El numero lo dibujamos en el eje Y un poquitin para abajo, osea en la fila de pixeles 35 de la ventana
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

	// Si la bola en el eje X su centro es menor a 0 eso quiere decir que pasamos el borde izquierdo de la ventana
	// por lo tanto rightPaddle anoto punto, por tanto se incrementa +1 su score
	if ball.pos.x < 0 {
		rightPaddle.score++
		// Y restauramos la posicion de la bola en el centro para iniciar nueva partida
		ball.pos.x = float32(winWidth) / 2
		ball.pos.y = float32(winHeight) / 2
		// Y colocamos el estado del juego en start, ¿Porque? para justamente
		// que el juego este en pausa y si le damos a espacio arrancamos el juego
		state = start
	}

	// Ahora en el eje X si el centro de la pelota es mayor al ancho de la ventana eso quiere decir que la pelota esta pasando
	// el borde derecho de la ventana, por lo tanto leftPaddle anoto punto
	if ball.pos.x > float32(winWidth) {
		leftPaddle.score++
		ball.pos.x = float32(winWidth) / 2
		ball.pos.y = float32(winHeight) / 2
		// Lo mismo aca con el estado del juego, si el leftPaddle marca score entonces el estado
		// del juego pasa a start, esto es para que justamente al darle a la tecla ESPACIO arrancar la nueva partida
		state = start

	}
	// Aca ahora vamos a modificar, lo que teniamos antes es que la pelota rebota en la paleta pero si el pixel central
	// de la pelota coincide con el eje vertical de la cara interna de la paleta, lo que vamos a cambiar
	// logicamente es que si el BORDE de la pelota toca la paleta entonces rebote, por tanto tenemos que tener en cuenta el radio de la pelota
	if ball.pos.x-float32(ball.radius) < leftPaddle.pos.x+leftPaddle.w/2 {
		if ball.pos.y > leftPaddle.pos.y-leftPaddle.h/2 && ball.pos.y < leftPaddle.pos.y+leftPaddle.h/2 {
			ball.xv = -ball.xv
			// Ahora bien algo que vamos a explicar y que quiza es un poco dificil es que hay veces que si reproducimos el juego
			// y la pelota choca con la paleta izquierda (leftPaddle) justo en ese microsegundo de choque hay veces que la pelota queda bugeada y rebota
			// muchisimas veces en el borde de la paleta ¿Esto porque sucede? La pelota cuando va viajando hacia la paleta
			// con una velocidad en X de -5 entonces choca la paleta y en el sig. fotograma la pelota adquiere una velocidad de +5 en el eje X
			// PEEEEROOO el juego en el fondo se va actualizando por FOTOGRAMAS, osea cada 16 microsegundos se genera un fotograma,
			// por lo que puede suceder que el borde de la pelota toque el borde de la paleta con una velocidad en X de -5 pero al hacer contacto
			// los bordes el fotograma no se actualice y por tanto la pelotita se meta un poquitin mas adentro de la paleta con la velocidad en X de -5
			// por tanto en el sig. fotograma la pelota efectivamente va a rebotar con una velocidad en X de +5 pero toca de nuevo el borde de la paleta ya
			// que, como dijimos, la pelota se metio un poquitito dentro de la paleta. Como consecuencia de todo esto es que la pelota
			// rebote muchisimas veces en el borde de la paleta

			// Una solucion a esto seria obviamente que se genere un fotograma cada -infinito microsegundos pero fisicamente es imposible
			// por tanto una opcion seria que se genere un fotograma con el minimo valor posible de fotograma generado por microsegundos en la funcion Delay() en main()
			// pero esto no lo podemos hacer porque el juego no correria bien para computadoras viejas
			// Entonces otra alternativa es decir que cuando la pelotita cambia el sentido de la velocidad en eje X
			// entonces tambien hacer que se actualice su valor en coordenada X, y que este valor del centro de la pelotita en coordenada X
			// se reposicione en el eje X del centro de la paleta + el ancho de la paleta izquierda/2 + el radio de la pelotita
			// con esta actualizacion del centro de la pelota en coordenada X nos aseguramos que la pelota cuando cambie el sentido de la velocidad automaticamente
			// el centro en X de la pelotita este por fuera de la paleta en el sig. fotograma, osea...
			ball.pos.x = leftPaddle.pos.x + float32(leftPaddle.w)/2.0 + float32(ball.radius)
		}
	}

	// Aca sucede lo mismo pero de la paleta del lado derecho
	if ball.pos.x+float32(ball.radius) > rightPaddle.pos.x-rightPaddle.w/2 {
		if ball.pos.y > rightPaddle.pos.y-rightPaddle.h/2 && ball.pos.y < rightPaddle.pos.y+rightPaddle.h/2 {
			ball.xv = -ball.xv
			// Aca hacemos lo mismo que el choclo que explicamos antes pero para la paleta del lado derecho
			ball.pos.x = rightPaddle.pos.x - float32(rightPaddle.w)/2.0 - float32(ball.radius)
		}
	}
}

func (paddle *paddle) update(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		// Tenemos que verificar que no nos pasemos del margen superior de la ventana sino no podriamos dibujar
		if paddle.pos.y-float32(paddle.h)/2 >= 0 {
			paddle.pos.y -= 7
		}
	}

	if keyState[sdl.SCANCODE_DOWN] != 0 {
		// Tenemos que verificar que no nos pasemos del margen inferior de la ventana asi no dibujamos la paleta
		// pasando este margen
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

// Aca lo que haremos sera algo nuevo, vamos a aplicar una ENUMERACION en GO, en realidad usaremos un simulador de enumeraciones
// que tiene GO, es decir se pueden simular enumeraciones con constantes e IOTA

// ¿Que es esto de usar constantes e iota? 'IOTA' es un identificador reservado de Go que se puede usar en declaraciones de
// constantes y que se incrementa por cada uso dentro de una constante de bloque, por ejemplo
// yo en golang lo que puedo hacer es...

//		const (
//			Sunday = iota
//			Monday
//			Tuesday
//	        Wednesday
//			...
//		)

// Entonces lo que  sucede es que iota se inicializa en 0, y en cada linea va incrementando +1

// Entonces lo que haremos sera declara un nuevo tipo de dato llamado 'gameState' que en el fondo es un simple INT
// donde basicamente nos dice el estado del juego, es decir, si el juego esta en 'start' o 'play'
type gameState int

// Ahora lo que haremos sera crear una simulacion de enum en C, donde declararemos constantes, en realidad solo 2 constantes, 'start' y 'play' que seran
// de tipo gameState

const (
	start gameState = iota
	play
)

// Y ademas declaramos y definimos un estado aparte 'state' que seria el estado del juego inicial
var state = start

// --------------------------------------------------------

// Aca crearemos la parte del PUNTAJE, para ello lo vamos a hacer a mano ya que es poco texto
// lo que haremos sera crear una matriz de bytes que representara el 0, 1 y 2
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

// Ahora hacemos la funcion que dibuja los numeros donde recibiremos la posicion donde dibujaremos dicho numero 'pos'
// ademas recibimos el color del numero que dibujaremos, luego otro parametro 'size' para poder escalar el numero
// el numero que dibujaremos 'num' y el buffer de pixeles donde dibujamos 'pixels'
func drawNumber(pos pos, color color, size int, num int, pixels []byte) {
	// Lo que vamos a hacer es situarnos en la esquina superior izquierda para empezara dibujar de izq a derecha de arriba a abajo
	// Entonces en la posicion X que nos dan en 'pos' le restamos el escalado que nos dan (osea size) multiplicado por 3 (porque la matriz de numeros es de 5x3) dividido 2
	// porque justamente si estamos en el eje X parados en el centro entonces tenemos que ir hacia la mitad para atras, osea...
	startX := int(pos.x) - (size*3)/2

	// Ahora ATENCION!, nosotros con startX y startY (el q definiremos luego) lo que buscamos es posicionarnos en la esquina superior izquierda pero de CADA CUADRADO DE LA MATRIZ
	// 'nums', es decir, supongamos que el usuario ingreso como parametro num=2 entonces claramente tenemos que dibujar el '2' por lo que la matriz 'nums' nos quedaria:

	//		_____________
	//		|_1_|_1_|_1_|
	//		|_0_|_0_|_1_|
	//		|_1_|_1_|_1_|		El numero 2
	//		|_1_|_0_|_0_|
	//		|_1_|_1_|_1_|

	// Ahora bien con cuadradito nos referimos a cada celda de esta ultima matriz, es decir, con startX y startY nos queremos posicionar en la esquina superior izquierda de CADA CUADRADITO
	// Entonces el usuario en 'pos' nos va a dar las coordenadas (X,Y) pero del pixel central que corresponde a la celda central donde vamos a dibujar el numero completo
	// Por tanto a la posicion X que nos pasa el usuario le restamos el parametro 'size' (escalado), que este 'size' indica en pixeles cual es el tamaño de cada
	// cuadradito, por ejemplo si el usuario ingresa size=20 quiere decir que cada cuadradito es de 20x20 pixeles, luego a 'size' lo multiplicamos por 3 (porque la matriz es de 5x3, es decir
	// al multiplicarlos por 3 obtenemos todo el ancho de la matriz entera) y todo lo dividimos por 2 para justamente que startX se situe en el borde izquierdo de la matriz entera, lo dividimos por 2 ya que le restamos la
	// mitad de la matriz entera

	// Ahora la logica con startY es la misma pero en la coordenada Y logicamente, la coordenada Y logicamente hace referencia al pixel central del cuadradito central de la matriz, por tanto le tenemos que restar el escalado que ingresa
	// el usuario (size, ya que size representa en pixeles el alto de los cuadraditos) multiplicado por 5 ya que la matriz de alto tiene 5 cuadraditos, con todo esto obtenemos el alto total de la matriz, por tanto la tenemos que dividir por 2 para obtener la mitad
	// de altura de la matriz. Con todo esto conseguimos que startY se ubique en el borde superior de la matriz
	startY := int(pos.y) - (size*5)/2

	// Por tanto, hasta ahora lo que obtenemos es que el pixel (startX, startY) sea el pixel superior izquierdo del cuadradito superior izquierdo de la matriz

	// en la variable global 'nums', entonces con el valor 'num' que nos da el usuario seleccionamos el numero en la matriz 'nums' para iterar en ella, es decir...
	for i, v := range nums[num] {
		// Recordar que v es el valor y el i es el indice
		// Entonces si v=1 es logico que tenemos que pintar de blanco todo ese cuadradito de 20x20 pixeles (suponiendo que el usuario ingreso size=20)
		if v == 1 {
			// Ahora si iteramos por todo ese cuadradito
			for y := startY; y < startY+size; y++ {
				for x := startX; x < startX+size; x++ {
					// Ahora si pintamos de blanco los pixeles
					setPixel(x, y, color, pixels)
				}
			}
		}

		// Ahora bien, resulta que nosotros nos posicionamos en la esquina superior izquierda pero de cada CUADRADO!
		// es decir, de cada numerito de la matriz de numeros 'nums'
		// Lo que debemos hacer es que cuando salimos de iterar todo ese cuadradito ahora a startX sumarle + size
		// esto es para pasar en el eje X al sig. cuadradito y evaluar en la prox iteracion si hay que pintarlo o no
		startX += size

		// Ademas, una vez que ya hayamos dibujado 3 cuadraditos de una fila ahora debemos pasar al primer cuadrado pero de la fila de abajo
		// osea...
		if (i+1)%3 == 0 {
			startY += size // Aumentamos en size startY para pasar a la fila de abajo de la matriz de cuadraditos
			// Y posicionamos en eje X al primer cuadradito de la matriz de la nueva fila de abajo
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
	// Vamos a dibujar la bola en la mitad de la ventana
	ball := ball{pos{float32(winWidth) / 2, float32(winHeight) / 2}, 20, 7, 7, color{255, 255, 255}}

	keyState := sdl.GetKeyboardState()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {

			case *sdl.QuitEvent:
				return
			}
		}

		// Si el estado del juego esta en 'play' quiere decir que estamos jugando, por lo tanto
		// tenemos que ir actualizando todo
		if state == play {
			// Aca dibujamos el numero '2' de prueba a ver que tal funciona la funcion drawNumber(), vamos a dibujar el numero
			// 2 en el centro de la ventana
			// drawNumber(pos{float32(winWidth) / 2, float32(winHeight) / 2}, color{255, 255, 255}, 20, 2, pixels)
			player1.update(keyState)
			player2.aiUpdate(&ball)
			ball.update(&player1, &player2)

		} else if state == start { // Si el juego esta en estado de start
			// Aca lo que haremos sera consultar en el arreglo keyState por la tecla ESPACO, si vale distinto a 0 quiere decir que esta presionada
			if keyState[sdl.SCANCODE_SPACE] != 0 {
				// Si la tecla espacio esta presionada entonces cambiaremos el estado del juego de start a play nuevamente para justamente iniciar una nueva partida
				// pero tambien debemos resetear los puntajes de los usuarios
				// Pero ojo, para resetear el scoe tambien debemos verificar que el puntaje de alguno de los 2 usuarios este en 3 que es el puntaje
				// maximo
				if player1.score == 3 || player2.score == 3 {
					player1.score = 0
					player2.score = 0
				}
				state = play
			}
		}

		clear(pixels)

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
