// Vamos a ver los paquetes bufio y os para que sirven

// El paquete OS permite interactuar con el SO, por ejemplo gracias a este podemos manipular
// los FD (File Descriptors), lectura, escritura, manipulacion de variables de entorno, gestion de procesos, etc

// Por otro lado el paquete BUFIO nos permite realizar operaciones de entrada/salida bufferizadas

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Lo que haremos sera en el STDIN (FD=0) leer de ella, osea aca usaremos el paquete OS
	// y lo que haremos sera scannear lo que el usuario ingrese en el archivo STDIN y guardarlo en la variable 'scanner'
	scanner := bufio.NewScanner(os.Stdin)

	// Ahora lo que va a implementar la computadora es una BUSQUEDA BINARIA, que recordemos que se basa en mitades, es decir, si la computadora
	// tiene que adivinar un numero entre 1 y 100 entonces dira 50, si el usuario le dice que el 50 es muy grande entonces dira un numero entre el 1 y 50, osea 25, y asi...
	// Es decir, los numeros que dice la computadora se basan en mitades de intervalos

	// Entonces definimos 2 variables que es el minimo que tiene que decir el usuario y el maximo
	low := 1
	high := 100

	// Vamos a imprimir una linea diciendole al usuario que piense un numero entre 1 y 100
	fmt.Println("Piensa un numero entre el", low, "y el", high)
	// Ahora le decimos al usuario que presione ENTER cuando termino de colocar el numero
	fmt.Println("Presione ENTER")
	// Ahora vamos a scannear lo que puso el usuario en STDIN
	scanner.Scan()

	for {
		// Ahora vamos a arrancar por la mitad del intervalo entre 1 y 100, osea 50, este numero dira primeero la computadora
		guess := (low + high) / 2
		// Ahora le comunicamos al usuario sobre que numero eligio la computadora
		fmt.Println("Creo que el numero es:", guess)
		// Ahora le preguntamos al usuario si el numero que eligio el es mas chico o mas grande que el que dice la computadora
		// o si adivino
		fmt.Println("El numero es?:")

		fmt.Println("(a) muy grande")
		fmt.Println("(b) muy chico")
		fmt.Println("(c) es el correcto")

		// Ahora scanneamos nuevamente lo que ingrese el usuario sobre que respondera
		scanner.Scan()
		// Lo que ingreso el usuario en el STDIN lo guardamos en la variable siguiente usando la funcion Text() que lo que hace
		// es que el ultimo texto ingresado en el scanner el Text() lo toma y lo guarda como una string
		respuesta := scanner.Text()

		// Ahora atendemos cada caso, si el usuario ingreso (a), (b) o (c)
		switch respuesta {
		case "a":
			// Si el usuario ingreso 'a' eso quiere decir que el numero que dijo la computadora
			// es muy grande, por lo tanto a hight sera lo que dijo la computadora - 1 y asi implementar la busqueda binaria
			high = guess - 1
		case "b":
			// Si entramos a este caso es porque el numero que dijo la computadora es muy chico, por tanto
			// seria el caso contrario, el minimo es lo que dijo la computadora +1
			low = guess + 1
		case "c":
			// Si estamos en este caso es porque la computadora adivino
			fmt.Println("Correcto")
			// Por ultimo salimos del ciclo for ya que es un loop eterno
			return
		default:
			// Si el usuario ingresa cualquier cosa que no sea (a) (b) o (c) le diremos al usuario
			// que ingeso un comando invalido
			fmt.Println("Comando invalido, intente de nuevo")
		}

	}
}
