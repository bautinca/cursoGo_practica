package main

import (
	"fmt"
	"time"
)

// Recordas cuando hablabamos que los canales son simples? Es decir, solo reciben un dato y se puede leer de ellos un dato

// Bueno ahora vamos a explicar como hacer que los canales reciban en el mismo tiempo varios datos, a estos se les llama BUFFERED CHANNELS

// Crearemos la funcion que lee datos del channel ch
func imprimir(in <-chan string) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

// Ahora la funcion enviarMensaje que lo que hace es escribir un dato en el canal ch
func enviarMensaje(out chan<- string, numero int) {
	for {
		out <- fmt.Sprintf("Mensaje: %d", numero)     // Guardamos un string en el canal out
		fmt.Println("Enviando mensaje func:", numero) // Imprimimos que guardamos el mensaje en el canal out
	}
}

func main() {

	// Vamos a crear un BUFFERED CHANNEL, para ello tambien se hace con make() pero en el segundo parametro le indicamos cuantos
	// datos queremos que reciba
	ch := make(chan string, 5) // El channel 'ch' puede recibir simultaneamente 5 datos

	// Entonces que pasaria si al channel ch le enviamos solo 2 datos? Bueno podemos escribir otros 3 datos mas en el o leer esos 2 datos
	// que tiene el canal

	// Si al canal le mandamos simultaneamente 5 datos entonces esta bloqueado, por lo tanto no se pueden escribir mas datos en el por lo tanto
	// solo sera de lectura

	// Lo que haremos ahora sera escribirle simultaneamente 4 datos al canal en el propio main
	for i := 1; i < 5; i++ {
		go enviarMensaje(ch, i) // Entonces lo que sucederia es que se enviar los 4 mensajes al mismo tiempo es decir, se generarian 4 gorutinas distintas
	}

	imprimir(ch) // Aca imprimira si o si todos los datos del channel ya que no es una gorutina por lo tanto
	// antes de que finalice el main si o si se tiene que terminar de ejecutar el imprimir()
}

// Lo que sucederia es que el channel ch NUNCA SE BLOQUEA ya que tiene una capacidad de 5 y yo constantemente le estoy pasando 4 datos, y cuando leo los datos con el imprimir()
// lo estoy vaciando nuevamente, es por eso que nunca llega al estado de bloqueo el canal

// Ahora vamos al channels7.go
