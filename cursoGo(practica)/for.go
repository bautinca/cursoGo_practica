package main

import "fmt"

func main() {

	// Al principio habiamos dicho que GO es un lenguaje ESTRUCTURADO, es decir, soporta estructuras de control, como
	// por ejemplo la gran mayoria de los lenguajes, el FOR, la sintaxis del FOR es pasarle un condicional y mientras esta condicion
	// se cumpla siempre se va a ejecutar el bloque de codigo que le pasemos, osea:

	contador := 0
	for contador < 4 {
		fmt.Println("Ejecucion")
		contador++
	}

	// En este ultimo ejemplo se imprimiria en el STDOUT el 'Ejecucion' 4 veces, es decir, es casi
	// identico al while

	// Tambien otra manera de correr el for es identicamente a C, osea:

	for i := 0; i < contador; i++ {
		fmt.Println(i)
	}
}

// Ahora vamos al archivo if.go
