package main

import (
	"fmt"
	"time"
)

// Vamos a hacer una optimizacion del programa anterior en channels3.go, lo molesto que teniamos antes
// era que teniamos que estar chequeando constantemente con el 'ok' si nos escapabamos del channel

// Entonces para evitar esta verificacion lo que podemos hacer es es usar el RANGE! ya que el range puede recorrer el channel y justo terminar de recorrerlo cuando no tenemos valor en el channel
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
		close(numero)
	}()

	go func() {
		// Aca lo que haremos sera usar el RANGE para recorrer el channel 'numero' por tanto no hace falta verificar si leemos afuera del channel
		// ya que el RANGE se ajusta al channel en cuestion
		for x := range numero {
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	// Ahora en el main hacemos lo mismo para recorrer el channel 'cuadrado' usando el RANGE
	for x := range cuadrado {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

// Ahora vamos a channels5.go
