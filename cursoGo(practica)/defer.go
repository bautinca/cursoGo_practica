package main

import "fmt"

func primera() {
	fmt.Println("Primera")
}

func segunda() {
	fmt.Println("Segunda")
}

func main() {
	// Vamos a ver DEFER, se utiliza para POSPONER la ejecucion de una funcion hasta que la funcion
	// que la contiene termine de ejecutarse.
	// Esto significa que cualquier funcion que se haya deferido (defer) se ejecutara justo antes de que la funcion que la contiene
	// ejecure su return

	// Un ejemplo muy simple es justamente aca dentro del main()
	// Podemos deferir el PrintLn() que justamente es una funcion
	defer fmt.Println("Esto se ejecutara despues")
	fmt.Println("Esto se ejecutara primero")

	// Entonces lo que suceder es que primero veremos el "Esto se ejecutara primero"
	// Y antes de que finalice el main() se ejecuta el defer, osea "Esto se ejecutara despues"

	fmt.Println("---------------------------")

	// Ahora veamos otro ejemplo, aca usaremos las funciones primera() y segunda()
	defer primera()
	segunda()

	// Aca en este ultimo ejemplo sucedera lo mismo que antes, es decir, primero se ejecutara segunda() y al finalizar el main primera()
	// Pero como antes tenemos otro defer entonces al finalizar el main() se ejecutaran todos los defer
}

// Ahora vamos al archivo defer2.go
