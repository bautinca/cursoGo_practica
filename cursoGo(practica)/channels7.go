package main

import (
	"fmt"
	"time"
)

// Vamos a ver por ultimo la instruccion SELECT, es una estructura de control que se usa para manejar multiples operaciones de canales de manera
//concurrente. Nos permite seleccionar entre multiples casos de operaciones de canal que esten listo para ser procesados.

// Es casi igual al SWITCH pero especificamente diseñado para CANALES. Entonces como funciona igual que el SWITCH tambien vamos a tener CASES

// Veamos un ejemplo

func main() {

	// Creamos 2 canales
	canal1 := make(chan int)
	canal2 := make(chan int)

	// Ahora creamos una funcion anonima que se ejecuta como una gorutina que lo que hace
	// es que tendra un sleep de 2 segundos donde al canal1 se le escribe el entero 1
	go func() {
		time.Sleep(2 * time.Second)
		canal1 <- 1
	}()

	// Por otro lado otra funcion anonima que hace lo mismo que la anterior pero le escribe el entero 2
	// al canal canal2, pero tiene un sleep de 1 segundo
	go func() {
		time.Sleep(1 * time.Second)
		canal2 <- 2
	}()

	// Entonces lo que sucederia es que primero al canal2 se le agrega el 2 y luego al canal1 se le agrega el 1
	// esto es porque el Sleep cuando el canal2 recibe el 2 es menor (1 segundo) que cuando el canal1 recibe el 1 (2 segundo)

	// Ahora activamos el SELECT, lo que hace el esquema valor := <- canal, es decir,
	// El canal espera a recibir un valor y lo asigna a la variable 'valor'

	select {
	case num := <-canal1:
		fmt.Println("Recibido de canal1:", num)

	case num := <-canal2:
		fmt.Println("Recibido de canal2:", num)
	}

	// Es decir, se ejecuta o un case o el otro, entonces si recibimos primero un valor de canal1 entonces ahora num le corresponde ese valor
	// e imprimimos cual es ese valor con el println()

	// Por otro lado si canal2 es el que primero recibe el valor entonces se ejecuta el segundo case

	// Siguiendo la logica el que primero recibe el valor es canal2 por tanto se tendria que ejecutar el segundo case y finalmente salir del main()

}
