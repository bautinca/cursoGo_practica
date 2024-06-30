package main

import (
	"fmt"
	"time"
)

// Vamos a usar el mismo codigo que teniamos antes en channels4.go, quiza el error que teniamos antes es que
// quedaba medio feo tener funciones gorutinas anonimas dentro del main(), queda un poco desprolijo

// Entonces lo que vamos a hacer es hacer que dejen de ser anonimas para declararlas y definirlas fuera del main

// Primero la funcion que le mete numeros al canal 'numero'
func generarNumeros(out chan<- int) { // Atencion aca, vemos algo nuevo que es el '<-' luego del tipo de dato 'chan', esto significa
	// que esta funcion efectivamente recibe un canal, pero si le colocamos un '<-' luego del tipo de dato 'chan' indica que el canal que recibe es SOLO PARA ESCRIBIRLE DATOS
	// Es decir, NO SE PUEDEN LEER DATOS DEL CANAL QUE RECIBE
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

// Ahora la funcion que toma los numeros del canal 'numero' y los eleva al cuadrado para guardarlos en el canal 'cuadrado'
func elevarAlCuadrado(in <-chan int, out chan<- int) { // Aca es la misma logica que antes, estamos diciendo que el canal 'in' UNICAMENTE SERVIRA PARA
	// LEER DATOS DE EL, en cambio el canal 'out' SOLO SIRVE PARA ESCRIBIR DATOS EN EL
	for x := range in {
		out <- x * x
	}
	close(out)
}

// Por ultimo la funcion que imprime tomando valores del canal 'cuadrado'
func imprimir(in <-chan int) { // Al poner '<-' delante del 'chan' decimos que esta funcion imprimir() recibe un canal en el cual SOLO
	// VAMOS A PODER LEER DE EL
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

// Colocando los '<-' delante o despues del tipo de dato 'char' le estamos diciendo a GO de que tipo son los canales que recibe cada funcion
// si de lecutra o de escritura, entonces GO va a estar mas al al tanto de las cosas, por ejemplo si en un canal que escribo datos accidentalmente accedo para leer en el
// entonces EN TIEMPO DE COMPILACION GO	me lanza una advertencia, y es mucho mas eficiente esto que que el error se me genere en tiempo de ejecucion

// Ahora si la funcion main va a quedar mucho mas prolija
func main() {

	// Creamos las 2 canales
	numero := make(chan int)
	cuadrado := make(chan int)

	// Funcion que coloca numeros en el canal 'numero'
	go generarNumeros(numero)

	// Funcion que toma los numeros del canal 'numero', los eleva al cuadrado y los guarda
	// en el canal 'cuadrado'
	go elevarAlCuadrado(numero, cuadrado)

	// Funcion que toma los numeros del canal 'cuadrado' y los imprime
	go imprimir(cuadrado)

	// Si aca no colocamos nada entonces al ejecutar el programa no se generaria nada ya que la funcion main() va a llegar al final
	// y por lo tanto las gorutinas no se van a ejecutar, para dejarlo en pausa o suspendido colocaremos un scanf
	var input string
	fmt.Scanln(&input)

}

// Ahora vamos a channels6.go
