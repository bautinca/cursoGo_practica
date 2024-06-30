package main

import "fmt"

func main() {
	// Bueno, vamos a hablar del IF, aca funciona tal cual igual que en casi todos los lenguajes, seria
	// bastante al pedo complicarsela

	// Lo que vamos a hacer es combinar el for con el if
	contador := 0
	for contador <= 5 {

		// Aca colocamos el if, como vemos funciona igual que en todos los lenguajes, un condicional y luego si da true entonces
		// colocamos un bloque de codigo para que ejecute
		if contador == 1 || contador == 4 {
			fmt.Printf("El contador vale 1 o 4\n")
		} else { // Y aca logicamente el ELSE! por si no entramos al IF
			fmt.Println("El contador no es ni 1 ni 4")
		}
		contador++
	}

	// Tambien una novedad del IF en GO es que podemos declarar y definir variables en la misma linea del IF, obviamente el scope de lo que
	// declaramos y definimos en el IF vive solo dentro del IF
	if nombre := "Juan"; len(nombre) > 5 {
		fmt.Println("El nombre tiene mas de 5 letras")
	} else {
		fmt.Println("El nombre tiene 5 letras o menos")
	}

}

// Ahora vamos al archivo switch.go
