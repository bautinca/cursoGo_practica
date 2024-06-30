package main

import (
	"fmt"
	"time"
)

// Ahora lo que vamos a hacer es tal cual lo mismo que antes, pero en el ejercicio anterior era infinito y no terminaba
// mas el main(), lo que vamos a hacer ahora es que ahora el envio de datos en un channel tenga un tope

// La funcion main
func main() {
	// Creamos los 2 canales
	numero := make(chan int)
	cuadrado := make(chan int)

	// Ahora declaramos una gorutina anonima
	go func() {
		for x := 0; x < 5; x++ {
			numero <- x
		}

		// Una vez que ya escribimos los 5 datos en el channel 'numero' lo que vamos a hacer es que se cierre el canal ya que no vamos
		// a escribir mas en dicho canal
		// NOTA: Recordar que no es que escribimos los 5 datos de una en el canal, sino que recordar que se escribe 1 dato y se bloquea hasta que otra gorutina reciba el dato
		// y recien cuando recibe el dato luego se escribira otro dato en el canal 'numero'
		// Entonces cuando ya se escribieron 5 datos (0 1 2 3 4) en el canal lo cerramos ya que no escribiremos mas en el
		close(numero)
	}()

	// Ahora la gorutina que recibe los datos del channel 'numero' y les saca el cuadrado para guardarlo en el channel 'cuadrado'
	go func() {
		for {
			x, ok := <-numero // El segundo variable es 'ok' ya que puede pasar que leamos un dato que no esta en el channel entonces cuando sucede eso
			// es porque ya leimos todo del channel 'numero' por tanto salimos del for
			if !ok {
				break
			}
			cuadrado <- x * x
		}
		// Cuando ya guardamos todos los datos en el canal 'cuadrado' lo cerramos
		// ya que no escribiremos mas datos en el canal
		close(cuadrado)
	}()

	// Ahora en el main creamos el for true para ir tomando lo datos del channel 'cuadrado' e imprimirlos
	for {
		x, ok := <-cuadrado
		if !ok { // Si ok=false quiere decir que ya terminamos de leer los datos del channelk 'cuadrado' por lo tanto salimos y dejamos de imprimir
			// Si no fuera por este ok entonces se imprimiria en la stdout infinitamente valores 0's ya que justamente nos pasamos del channel
			break
		}

		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

	// Finalmente cerramos el channel 'cuadrado' ya que no leeremos mas datos de el
}

// Ahora vamos a channels4.go
