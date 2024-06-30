package main

import "fmt"

// Ahora vamos a ver PANIC y RECOVER

// PANIC es un error que detiene la ejecucion del programa simplemente eso,
// osea detiene la ejecucion de todo el programa

// Y el RECOVER sirve para justamente RECUPERAR el mensaje de error que envia el PANIC, para ello
// vamos a hacer un ejemplo

// Primero vamos a hacer una funcion que recupera el mensaje de error del PANIC
func funcDeRecu() {
	// Si se ejecuta este if es porque efectivamente tenemos un recover que hacer ya que r da distinto de nulo y es el mensaje de error del panic
	if r := recover(); r != nil {
		fmt.Println("Se recupero de un panic:", r)
	}
}

// Ahora la funcion que emite el PANIC
func funcDePanic() {
	// Aca colocaremos un DEFER ¿Porque? Porque si llamamos a la funcion funcDeRecu despues del PANIC
	// entonces esa linea nunca se ejecutaria ya que lo que hace el PANIC es detener absolutamente toda la ejecucion
	defer funcDeRecu()
	panic("Esto es un error critico!")

	// Entonces asi lo que logramos es que la funcDeRecu() se tenga en cuenta y DESPUES ponemos el PANIC,
}

// Ahora la funcion main para probar las funciones
func main() {
	funcDePanic() // Entonces lo que haria esta funcion es causar un PANIC y es por ello que dentro de ella
	// llamamos a la otra funcion funcDeRecu() que hace un RECOVER para tomar el mensaje de error del PANIC que se emitio e imprimirlo

	// Por otro lado ahora hacemos un nuevo print para comprobar que si al mensaje de PANIC lo recuperamos entonces
	// la ejecucion del programa sigue
	fmt.Println("Esto no se deberia imprimir")

	// Ahora bien, es importante decir que cuando un PANIC no tiene un RECOVER entonces la ejecucion del programa es COMPLETA, es una detencion abrupta en la
	// ejecucion del programa
}

// Ahora vamos al archivo punteros.go
