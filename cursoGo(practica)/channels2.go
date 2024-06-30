package main

import (
	"fmt"
	"time"
)

// Veamos un poco mas sobre los CHANNELS

func main() {

	// Primero vamos a declarar 2 canales simples de INTs
	numero := make(chan int)
	cuadrado := make(chan int)

	// Lo que haremos sera crear 2 gorutinas pero de funciones anonimas,
	// Lo que hara la primera gorutina es ir escribiendo de a +1 datos en el canal 'numero', osea escribe 0 luego 1, luego 2, y asi...

	go func() {
		for x := 0; ; x++ { // Este es un for comun y corriente, lo unico es que no le pasamos ninguan condicion, lo dejamos vacio
			numero <- x
		}
	}()

	// Por otro lado otra gorutina que lo que hace es tomar el dato del channel 'numero', a ese dato multiplicarlo por si mismo e instar
	// el resultado en el channel 'cuadrado'

	go func() {
		for { // while true
			x := <-numero     // Tomamos el dato del canal 'numero'
			cuadrado <- x * x // A ese dato que tomamos lo operamos y lo metemos en el channel 'cuadrado'
		}
	}()

	// Ahora volvemos al main, entonces mientras se ejecuta todo lo anterior (osea las 2 gorutinas de arriba), lo que hara
	// el main es ir tomando dato por dato del channel 'cuadrado' e imprimirlo en pantalla, y colocaremos un sleep para que se vea bien el proceso
	// en como el main va tomando los datos del channel 'cuadrado'
	for {
		fmt.Println(<-cuadrado)
		time.Sleep(1 * time.Second)
	}

	// Si dejamos el main asi como esta entonces es logico que nunca va a terminar porque tenemos el for true anterior
	// entonces este programa nunca terminaria

}

// Ahora vamos a channels3.go
