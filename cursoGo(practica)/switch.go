package main

import "fmt"

func main() {
	// Vamos a ver la estructura de control SWITCH, que su funcionamiento es igual que en
	// la de C

	dia := 3 // Declaramos y definimos la variable, recordar que con := la declaracion es automatica

	// Ahora vamos a hacer un switcheo de la variable 'dia'
	switch dia {
	// Si dia=1...
	case 1:
		fmt.Println("Lunes")

		// Si dia=2...
	case 2:
		fmt.Println("Martes")

		// Si dia=3...
	case 3:
		fmt.Println("Miercoles")

		// Si dia=4...
	case 4:
		fmt.Println("Jueves")

	case 5:
		fmt.Println("Viernes")

		// El default se ejecutaria si dia es cualquier otro valor
	default:
		fmt.Println("Fin de semana")
	}

	// Pero tambien en los 'case' podemos colocar condiciones y hacer de cuenta como si fueran muchos IFs
	// Y aca algo a destacar es que en el switch no hay que acarar que variable vamos a switchear ya en los cases hacemos alusion
	// directamente a condiciones, es decir, literalmente es como si fueran muchos IFs los cases
	hora := 12

	switch {
	case hora < 12:
		fmt.Println("Buenos Dias")

	case hora >= 12 && hora < 18:
		fmt.Println("Buenas tardes")

	default:
		fmt.Println("Buenas noches")

	}

	// Tambien, al igual que con el IF, en la misma linea del switch podemos declarar variables
	switch animal := "Ave"; animal {
	case "Elefante":
		fmt.Println("El animal es un", animal)

	case "Ave":
		fmt.Println(fmt.Sprintf("El animal es un %s", animal))

	case "Leon":
		fmt.Println("El animal es un " + animal)

	case "Toro":
		fmt.Println("El animal es un", animal)
	}

	// Si nos fijamos, a diferencia de C, aca no es necesario colocar el BREAK! ya que cuando terminamos
	// de leer en el case indicado es obvio que no vamos a leer ningun case mas, es decir, de todos los cases siempre se ejecuta 1
	// por tanto es bastante obvio que siempre al final de un case hay un break

	// Tambien la novedad que tiene GO es que en el caso del switcheo agrego la instruccion FALLTHROUGH que es justamente lo contrario al BREAK
	// es decir, si queremos que justamente no se lea el break automatico y que se sigan evaluando cases, por ejemplo:

	var numero int = 10

	switch {

	case numero < 15:
		fmt.Println("El numero es menor a 15")
		fallthrough

	case numero < 11:
		fmt.Println("El numero es menor a 11")

	case numero > 5:
		fmt.Println("El numero es mayor a 5")

	}

	// En este ultimo caso entrariamos al case <15 y como dice fallthrough entonces
	// indica que se siga leyendo dentro del switch. Luego al seguir leyendo
	// tambien se ejecutaria el case <11, pero este case no tiene el fallthrough, por
	// tanto el case >5 no se ejecutaria si bien cumple con la condicion
}

// Ahora vamos a ir al archivo arrays.go
