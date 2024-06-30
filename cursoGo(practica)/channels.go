package main

import (
	"fmt"
	"time"
)

// Los channels en go son un mecanismo de comunicacion y sincronizacion entre gorutinas que permite transferencia
// de datos seguros entre ellas.

// Los channels, como bien dijimos, tambien son utiles para que haya sincronizacion entre la ejecucion de las gorutinas, por ejemplo
// podemos hacer que una gorutina espere a recibir un valor antes de continuar ejecutandose

// Podemos enviar tipos de datos de una gorutina a otra a traves de los canales de todo tipo, por ejemplo un INT, string, struct, etc.

// Los CHANNELS admiten 2 operaciones basicas o sintaxis nuevas que veremos:

//		<- : Envia los datos y recibir los datos

// La operacion de envio coloca justamente un dato en el canal mientras que la operacion de recibir recibe un dato del canal, ahora bien
// estas operaciones son BLOQUEANTES lo que significa que si una gorutina envia un dato automaticamente la gorutina que envia el dato se BLOQUEARA hasta que la otra gorutina reciba el dato, es decir
// una gorutina se bloqueara hasta que la otra realice la operacion complementaria

// Ejemplo:

// Esta funcion envia los datos al channel
func enviarDatos(c chan string) { // c seria el canal, logicamente es de tipo de dato chan
	// Lo que haremos para enviarle datos al channel es un for true pasandole el string "Ping"
	for true {
		c <- "Ping" // El operador '<-' luego del channel c, sirve para escribir un dato en el channel
	}
}

// Esta funcion toma los datos del channel y los imprime
func imprimirDatos(c chan string) {
	// Primero creamos una variable contador
	contador := 0
	// Ahora un for true (while true)
	for {
		contador++
		fmt.Printf("%s: %d\n", <-c, contador) // El operador <- sirve para leer un dato del channel c
		// Y colocamos un Sleep para que se imprime de a poco y poder ver el proceso
		time.Sleep(time.Second * 1) // 1 segundo por cada dato que leemos del channel
	}
}

// La pregunta es ¿Porque los colocamos dentro de un while true (for true)? Para que justamente ambas gorutinas se ejecuten constantemente y nunca salir de ellas
// Es decir, a medida que el imprimirDatos mete el 'Ping' en el channel siempre el enviarDatos() va a recibirlos e imprimir en pantalla junto con el valor del contador

func main() {
	// Para crear un canal se hace tambien con make() al igual que para crear slices o maps
	ch := make(chan string) // Aca creamos un canal de tipo STRING

	// Ahora vamos a crear una gorutina para enviarle STRINGs al channel
	go enviarDatos(ch)
	// Simultaneamente mientras que escribimos datos en el channel tambien los vamos a leer e imprimir:
	go imprimirDatos(ch)

	// Si corremos el programa asi como esta hasta ahora entonces finalizaria, es decir no funciona porque
	// recordemos que a medida que se ejecutan las gorutinas el programa principal sigue, por lo tanto llegamos al final del main() y termina la ejecucion

	// Para hacer que el main nunca termine la ejecucion lo que podemos hacer es dejarlo en pausa, esto es por ejemplo esperando a recibir un dato de entrada con el famoso
	// Scanln()
	var input string
	fmt.Scanln(&input) // Entonces al momento que escriba algo vamos a llegar al final del main() por lo tanto va a finalizar la ejecucion, obviamente las gorutinas van a terminar
}

// NOTA: Siempre que nosotros enviamos un valor al channel la otra contraparte esta ESPERANDO ESE VALOR para recibirlo y manipularlo como quiera
// Es decir, una gorutina envia el valor y queda suspendido hasta que la otra gorutina lo tome, una vez que la gorutina toma el valor, entonces la otra gorutina se va activar nuevamente
// y enviar otro valor al channel, y se queda nuevamente en pausa hasta que la otra gorutina tome el valor, y asi...

// Es decir, no es que una gorutina escribe 'x' valores en el channel y luego la otra toma los valores que se le cante, siempre la relacion es 1 a 1, es decir, se escribe 1 dato, se lee 1 dato,
// se escribe 1 dato y se lee 1 dato, y asi...

// Es decir, en resumen este simple channel que vimos TIENE LA CAPACIDAD DE 1 SOLO VALOR!

// Ahora vamos a channels2.go
