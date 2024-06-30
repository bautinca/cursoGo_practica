package main

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// Primero vamos a copiar todo lo que teniamos en sdl2.go y para crear el juego ping-pong vamos a crear
// codigo pero teniendo como base sdl2.go

// Logicamente en el juego ping-pong tenemos una pelota, por lo tanto sera un struct que describa caracteristicas
// de la bolita
type ball struct {
	pos    pos     // Seria la posicion de la bolita
	radius int     // El radio de la bolita
	xv     float32 // Velocidad de la bolita en eje x
	yv     float32 // Velocidad de la bolita en eje y
	color  color   // Color de la bolita
}

// Tambien vamos a hacer otro struct que sea la posicion
type pos struct {
	x, y float32 // Logicamente la pos son las coordenadas X e Y, colocamos como tipo de dato float32 ya que logicamente son datos mas precisos para
	// describir la pos. de los objetos del juego
}

// Que mas falta? Las 'raquetas' del ping pong, osea las paletas
type paddle struct {
	pos   pos   // Posicion de las paletas
	w     int   // Ancho de las paletas
	h     int   // Alto de las paletas
	color color // Color de las paletas
}

// Bien ahora vamos a hablar un poco en general como programar el juego, la idea principal seria en el main tener un bucle for donde
// dentro del bucle lo primero que obtendremos de todo son todas las entradas del juego por STDIN, es decir, tanto lo que escribe el usuario como lo que clickea
// Luego lo sig. que hacemos es que en base a la entrada actualizamos las cosas, por ejemplo si el usuario presiona 'W' entonces la paleta del juego se mueve para arriba
// y ese movimiento lo tenemos que graficar que esto seria el ultimo paso, una vez todo actualizado entonces dibujar!

// Osea la logica basica seria:

//		for {
//			1. Tomamos todas las entradas
//			2. Actualizamos
//			3. Dibujamos
// 		}

// Entonces lo primerisimo de todo que vamos a hacer es dibujar las paletas en la ventana
// Por lo tanto seria un metodo del struct paddle. Obviamente la funcion recibe 'pixels' que lo vimos en sdl2.go que es justamente para poder dibujar en el
// que representaria la matriz de toda la pantalla para dibujar
func (paddle *paddle) draw(pixels []byte) {
	// El paddle logicamente es un rectangulo, lo que vamos a hacer es que la posicion de la paleta represente el CENTRO
	// es decir si digo coordenadas (10,5) entonces esas coordenadas son del pixel central de la paleta
	// Entonces cuando dibujemos vamos a querer dibujar desde la parte superior izquierda de la paleta hasta la parte inferior derecha, osea todo el rectangulo
	// Entonces imaginandonos que las coordenadas de la paleta son en el centro entonces vamos a tomar la coordenada x de la paleta y restarle la mitad del ancho, esto es para
	// tomar el borde izquierdo de la paleta, osea:
	startX := int(paddle.pos.x) - paddle.w/2 // Es decir, a la pos central en el eje X de la paleta le restamos la mitad del ancho, esto es para pararnos en el borde izquierdo de la paleta en eje x
	// Ahora vamos a querer hacer lo mismo pero en el Eje Y, entonces en el eje Y nos tenemos que situar en el borde superior de la paleta
	// por lo tanto a la coordenada central de la paleta en el eje Y le restamos la mitad del alto de la paleta para situarnos en el borde superior de la paleta en eje Y
	startY := int(paddle.pos.y) - paddle.h/2

	// Listo, ya tenemos las coordenadas X e Y del borde superior izquierdo de la paleta que seria (startX, startY)
	// Ahora en base a este pixel lo que haremos sera dibujar toda la paleta de izquierda a derecha de arriba a abajo
	// obviamente tenemos que hacer un doble for, uno para el ancho de la paleta y otro para el alto de la paleta
	// Y luego cuando pasamos pixel por pixel para dibujarlo usamos la funcion setPixel() que implementamos en sdl2.go
	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels) // La paleta sera de color blanca
		}
	}
}

// Listo, ya tenemos el metodo que dibuja una paleta, ahora lo mismo pero con la pelotita
// obviamente seria un metodo del struct ball
func (ball *ball) draw(pixels []byte) {
	// Basicamente tenemos que dibujar un circulo, pero para hacer ello vamos a iterar sobre un cuadrado
	// inicialmente pero para cada pixel dibujado vamos a consultar si esta dentro del radio del circulo
	// Es decir, en resumen dibujamos un cuadrado comun y corriente pero colocamos un if en cada pixel si ese pixel esta dentro del radio
	// del circulo, obviamente estamos haciendo trabajo extra ya que estamos iterando sobre pixeles que no estan dentro del radio (las esquinas del cuadrado)
	// pero bueno, esta es una forma sencilla para implementarlo

	// Lo que haremos sera iterar por el radio negativo hasta radio positivo pero en el eje Y
	for y := -ball.radius; y < ball.radius; y++ {
		// Lo mismo para el eje X
		for x := -ball.radius; x < ball.radius; x++ {
			// Aca lo que hacemos es iterar por todo el cuadrado que forma el radio, pero logicamente las esquinas del cuadrado
			// no entrarian dentro del radio, por tanto lo que haremos sera colocar aca un condicional if si el pixel esta dentro del radio entonces lo pintamos de blanco
			// Lo que haremos sera sacar la distancia del pixel y el centro del circulo, y si esta distancia es menor al radio entonces el pixel lo tomamos
			// ¿Como sacar la distancia de un pixel hasta el centro del circulo? Es la norma de un vector!!, osea raiz(x*x + y*y), pero la raiz la pasamos del otro lado osea...
			if x*x+y*y < ball.radius*ball.radius {
				// Y aca aplicamos setPixel()
				setPixel(int(ball.pos.x)+x, int(ball.pos.y)+y, ball.color, pixels)
			}
		}
	}
}

// Haremos otro metodo de la bola pero que tenga que ver con el movimiento de la bola, es decir, seria como un actualizador de la pelotita
// para simular su movimiento. Ademas logicamente el metodo recibe 2 paletas para hacer que la pelotita cuando
// choque en estas paletas rebote y vaya en el otro sentido
func (ball *ball) update(leftPaddle *paddle, rightPaddle *paddle) {
	// La coordenada X de la pelotita se moveria segun su velocidad en el eje X osea...
	ball.pos.x += ball.xv
	// Lo mismo con la coordenada Y de la pelotita
	ball.pos.y += ball.yv

	// Ahora vamos a limitar para que la pelota no se escape de la ventana tanto en coordenadas X como en Y, si la pelota se escapa de la ventana
	// entonces la velocidad tiene que ser en el sentido contrario asi justamente la pelota rebota
	// Entonces si la pelota en coordenadas Y es menor a 0 eso quiere decir que la pelota se escapa en el margen inferior de la ventana por tanto la velocidad
	// en Y cambia de sentido
	// O tambien puede pasar que la pelota en la coordenada Y es mayor al alto de la ventana, eso quiere decir que la pelotita se esta escapando en el margen superior de la ventana, por tanto tambien cambiamos

	// pero ojo, a las coordenadas centrales de la pelotita le restamos el radio para el margen izquierdo y le sumamos el radio para el margen derecho
	// esto es justamente porque sino rebotaria la parte central de la pelotita y la otra mitad de la pelotita desaparece ya que se escapa de la ventana
	// Entonces para evitar esto y hacer que rebote la pelota a partir de los bordes de la pelotita tenemos que tener en cuenta su radio
	if int(ball.pos.y)-ball.radius < 0 || int(ball.pos.y)+ball.radius > winHeight {
		ball.yv = -ball.yv
	}

	// Si la pelotita se escapa por el margen izquierdo o derecho de la ventana eso quiere decir que el usuario contrario anoto un punto, por lo tanto la pelotita se resetea nuevamente a su lugar inicial
	// pero ojo, a las coordenadas centrales de la pelotita le restamos el radio para el margen izquierdo y le sumamos el radio para el margen derecho
	// esto es justamente porque sino rebotaria la parte central de la pelotita y la otra mitad de la pelotita desaparece ya que se escapa de la ventana
	// Entonces para evitar esto y hacer que rebote la pelota a partir de los bordes de la pelotita tenemos que tener en cuenta su radio
	if ball.pos.x < 0 || ball.pos.x > float32(winWidth) {
		ball.pos.x = 300 // Reseteamos el centro de la pelota en 300 en coordenada X
		ball.pos.y = 300 // Lo mismo que antes pero en eje Y
	}

	// Porque solo la pelota rebota en el margen inferior y superior? Porque asi es el juego PONG, es decir, logicamente la pelotita solo puede rebotar entre
	// el margen superior e inferior de la ventana

	// Ahora tenemos que hacer que la pelotita rebote si choca con los bordes de las paletas, tanto la del jugador como la del otro jugador
	// Entonces en el eje X si el pixel central de la pelotita es menor al pixel central de la paleta + la mitad del espesor de la paleta quiere
	// decir que la pelotita hizo contacto con el borde derecho del jugador1 que esta a la izquierda, por tanto el movimiento de la pelota tiene que cambiar de sentido
	if int(ball.pos.x) < int(leftPaddle.pos.x)+leftPaddle.w/2 {
		// Pero ojo ahora debemos chequear el eje Y donde justamente la pelota este dentro del intervalor en Y de la paleta, porque si la pelotita esta
		// por fuera de dicho intervalo eso quiere decir que la pelota sigue de largo, entonces vamos a asegurarnos que la pelotita esta en el intervalo Y de la paleta
		if int(ball.pos.y) > int(leftPaddle.pos.y)-leftPaddle.h/2 && int(ball.pos.y) < int(leftPaddle.pos.y)+leftPaddle.h/2 {
			ball.xv = -ball.xv
		}
	}

	// Lo mismo pero ahora con la paleta del lado derecho en eje X
	if int(ball.pos.x) > int(rightPaddle.pos.x)-rightPaddle.w/2 {
		if int(ball.pos.y) > int(rightPaddle.pos.y)-rightPaddle.h/2 && int(ball.pos.y) < int(rightPaddle.pos.y)+rightPaddle.h/2 {
			ball.xv = -ball.xv
		}
	}
}

// Ahora vamos a hacer el metodo update() pero de las paletas, las paletas logicamente pueden ir de arriba a abajo, es decir
// solo se mueven en el eje Y. Pero logicamente estas paletas se mueven de acuerdo al INPUT del usuario
// Entonces logicamete el update() de la paleta recibe como argumento el INPUT en teclado del usuario
func (paddle *paddle) update(keyState []uint8) {
	// Ahora algo que tiene SDL2 es que podemos scanear varias partes del teclado, por ejemplo
	// si colocamos el teclado la flecha hacia arriba seria la constante SCANCODE_UP, si el teclado esta presionado entonces vale distinto a 0
	// ya que un teclado no presionado vale 0
	if keyState[sdl.SCANCODE_UP] != 0 { // Con SCANCODE_UP seleccionamos el estado de la tecla UP, osea la felchita para arriba, si no esta presionada entonces vale 0
		// Aca la paleta se moveria hacia arriba, por lo tanto a la coordenada Y de la paleta
		// le restamos -5
		paddle.pos.y -= 5
	}

	// Ahora lo mismo pero si el usuario presiona la tecla de la flechita hacia abajo
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.pos.y += 5
	}
}

// Ahora lo que vamos a hacer es actualizar el movimiento de una segunda paleta que sera una inteligencia artificial!
// Entonces diremos que las coordenadas Y de esta paleta sean equivalentes a la coordenada Y de la pelotita, entonces con esto lo que logramos es que
// el centro de la paleta siempre esta alineada con el centro de la pelotita en Y
func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.pos.y = ball.pos.y
}

// Ahora una FUNCION que lo que hace es limpiar, clean(), es decir, la funcion update() anterior del paddle
// lo que hace es solamente pintar de mas la ventana de pixeles, es decir, si fuera solo por el update() y presionamos las flechas UP o DOWN seria
// como si estuvieramos alargando la paleta, esto es porque a medida que pintamos mas pixeles de blanco hacia arriba o abajo no restauramos a color negro los otros pixeles para
// dar esa sensacion de 'desplazamiento' de la paleta
func clear(pixels []byte) {
	// Entonces, lo que hara esta funcion sera absolutamente borrar toda la ventana, es decir, colocar todos los pixeles
	// de la ventana en negro
	for i := range pixels {
		pixels[i] = 0 // Esto es equivalente a poner el pixel en negro, es lo mismo que usar setPixel()
	}
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

	// Aca vamos a crear un objeto Paleta, donde el pixel central de la paleta estara en (100, 200) en la ventana
	// su ancho sera de 20 pixeles y su alto de 100 pixeles, y su color sera blanco...
	player1 := paddle{pos{100, 200}, 20, 100, color{255, 255, 255}}

	// Agregamos un segundo jugador, es decir, una segunda paleta centrada en la posicion contraria a la paleta1
	player2 := paddle{pos{700, 100}, 20, 100, color{255, 255, 255}}

	// Ahora la pelotita, en la ventana estara su pixel central en (300, 300) y tendra un radio de 20 pixeles, y su color blanco obviamente
	// Ademas hasta ahora la pelotita sera estatica, es decir, no tendra velocidad, por tanto xv=0 y yv=0
	// Recordar que el tamaño de la ventana en pixeles es de 800*600 y logicamente el centro esta en el pixel (400, 300)
	// Pero de todos modos la bola la dibujaremos un poco hacia la izquierda, la dibujaremos en las coordenadas (300, 300) donde estara el pixel central de la pelotita
	// ball := ball{pos{300, 300}, 20, 0, 0, color{255, 255, 255}}

	// Ahora la pelotita la hacemos con movimiento pero lentito, colocamos xv=2 y yv=2
	ball := ball{pos{300, 300}, 20, 2, 2, color{255, 255, 255}}

	// Ahora definimos KeyState que sera justamente lo que presione el usuario en teclado, osea del input del teclado
	keyState := sdl.GetKeyboardState() // Este keyState lo pasamos como argumento al metodo update() de la paleta paddle
	// Lo interesante del keyState es que es un ARREGLO de uint8!!, esto es porque es un arreglo donde se utiliza
	// para obtener el estado actual de TODAS LAS TECLAS DEL TECLADO, es por ello que es un arreglo

	// Ahora lo que vamos a hacer es un for ¿Porque? Para justamente tomar el INPUT del teclado o mouse, es decir, tomar eventos
	// que presione el usuario para ello hacemos un for True:
	for {
		// Ahora vamos a usar la funcion PollEvent() de SDL2, es crucial para manejar eventos dentro de un bucle de eventos
		// su proposito es consultar los eventos que se han acumulado en la cola de eventos y procesarlos, devuelve 1 si hay un eventos disponible en la cola y 0
		// si no hay ningun evento en la cola
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			// Ahora vamos a hacer un switch case para ver de que tipo es el evento
			switch event.(type) { // consultamos por el tipo de evento

			// Si el evento es para salir entonces finalizamos
			case *sdl.QuitEvent:
				return
			}
		}
		// Aca lo que haremos sera antes que nada limpiar toda la pantalla, es decir, para poner toda la pantalla en negro
		// Entonces en realidad lo que hacemos aca no es borrar todo lo que hay en la ventana, sino que lo que hacemos con este clear() es darle
		// esa sensacion de 'desplazamiento' a las paletas, entonces cada 16 microsegundos la pantalla se pondra toda en negro y se actualizara
		// los valores X e Y del centro de la paleta
		clear(pixels)

		// Ya teniendo los objetos paleta y pelota lo que haremos sera dibujarlos en la matriz 'pixels'
		player1.draw(pixels)
		ball.draw(pixels)
		player2.draw(pixels)

		// Ahora lo que hacemos es colocar un actualizador de la paleta, osea de player1, le colocamos el update() que definimos
		player1.update(keyState)
		// Y actualizador pero de la paleta2 que seria la inteligencia artificial que su centro siempre esta alineado con el centro de la pelotita
		// para ello usaremos el metodo de la paleta aiUpdate() para convertir a cualquier paleta en una ingeligencia artificial
		player2.aiUpdate(&ball)

		// Lo mismo pero con el actualizador de la pelotita y le pasamos los punteros a ambas paletas
		ball.update(&player1, &player2)

		pixelsPointer := unsafe.Pointer(&pixels[0])
		tex.Update(nil, pixelsPointer, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16) // Ahora este delay esta dentro de un for, por lo tanto el delay sera de muy poquititos microsegundos, ejemplo 16

	}
}

// Y listo, si ejecutamos todo esto asi como esta ya tenemos un prototipo de PONG funcional aunque con algunos errores minimos, por ejemplo
// que la pelota rebota en las paletas pero hasta que el centro de la pelota coincide con el eje en Y de las paletas pero del lado de la cara interna

// Otra de las cosas que tenemos que hacer es que en los juegos originales del PONG uno le podia dar 'efecto' al golpear la pelotita y asi dificultarle ganar al otro jugador
// es decir, si la pelotita golpeaba mas al borde de la paleta entonces la pelotita rebotaba pero con mayor apertura, en cambio si la pelotita choca mas hacia el centro de la paleta
// entonces la pelota rebota pero con menor apertura o amplitud

// Ademas otra de las cosas que debemos hacer es dibujar la puntuacion en la pantalla, es decir, si un jugador anoto un punto, el otro otro, y asi..
