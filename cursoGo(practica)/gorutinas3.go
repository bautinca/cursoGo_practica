package main

import (
	"fmt"
	"time"
)

// Vamos a implementar un programa que calcula la serie Fibonacci dado un numero usando Gorutinas
// Primero implementamos la funcion que calcula la serie de Fibonacci y como bien sabemos, sabemos que es recursiva
// Entonces a la funcion pasandole un numero indice me calcula la sucesion de Fibonacci, sabemos que la sucesion es 0,1,1,2,3,5,8,13,21, etc.
// Donde si por ejemplo le pasamos el indice 3 me tendria que devolver 2, si le pasamos 7 me tendria que devolver 13

func fibo(x int) int {
	if x < 2 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}

func animacion(retraso time.Duration) {
	for true { // Seria como un while true, NOTA: En GO no existe el while, ya que lo reemplaza el propio for
		fmt.Println("Procesando...")
		time.Sleep((retraso)) // Colocamos un sleep para ver como se va imprimiendo bien paso a paso el 'Procesando...'
		// si no fuera por este Sleep entonces se imprime a velocidad luz en cada linea el 'Procesando...'
	}
}

func main() {

	fmt.Println(fibo(3)) // Listo, perfectamente funciona la funcion fibo() que dado un indice me devuelve el valor en la sucesion fibonacci en ese indice

	// Ahora lo normal es que cuando le pasemos un indice muy alto el proceso tarde en ejecutarse calculando la sucesion de Fibonacci, por ejemplo:
	// fmt.Println(fibo(45)) // Al ejecutarse esta linea tarda aprox. 7 segundos en darnos el resultado

	// Entonces lo que vamos a hacer es que durante esos 7 segundos se este ejecutando otra cosa en otra gorutina! Para ello
	// Lo que haremos sera implementar una funcion que hace una animacion simulando a una 'carga' se llamara animacion()
	// Que lo que hace sera imprimir varias veces el texto 'Procesando...'

	// Lo que haremos sera  en una gorutina aparte generar la animacion
	go animacion(100 * time.Millisecond)
	// Ahora lo que sucedera es que en el fondo se esta ejecutando la gorutina animacion(), entonces mientras se ejecuta
	// lo que haremos sera que se piense el fibonacci de 45
	resultado := fibo(45)
	fmt.Printf("Fibonacci(45) es: %d\n", resultado)

	// Y listo! Mientras que el programa se piensa el Fibonacci de 45 se ejecuta siempre la linea 'Procesando...' entonces
	// cuando se termina de pensar el fibonacci de 45 llegamos al final del main() por lo tanto la gorutina finaliza automaticamente

}

// Ahora vamos al archivo channels.go
