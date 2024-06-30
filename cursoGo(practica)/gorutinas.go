package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Vamos a hablar sobre las CORUTINAS, como sabemos GO es un lenguaje CONCURRENTE y esto es una enorme
// ventaja

// Entones una GORUTINA se la puede definir como un hilo ligero manejado por GO, es decir, es al fin y al cabo un HILO como lo concemos
// en programacion

// Una GORUTINA es entonces una funcion que se ejecuta de manera CONCURRENTE con otras GORUTINAS dentro de un mismo programa o proceso
// Por lo que las GORUTINAS se deben aplicar si o si si lo que buscamos es concurrencia

// Como sabemos, si queremos q 2 funciones se ejecuten al mismo tiempo (en realidad no al mismo tiempo sino que se van turnando) estan los hilos donde
// justamente nos da esa 'sensacion' de que los 2 hilo se estan ejecutando paralelamente, osea al mismo tiempo

// Por ejemplo; C maneja hilos, pero porque se le agrego mucho mas adelante una libreria para poder manejar hilos, pero no es que
// C especialmente fue creado para concurrencia y el manejo de hilos.

// En cambio GO uno de sus fuertes es la concurrencia, ya que fue especialmente creado para la concurrencia y el uso de hilos

// Para entender bien esto de la concurrencia la antitesis de concurrencia es la EJECUCION SECUENCIAL

// Es decir, nosotros hasta el momento la ejecucion de los programas era SECUENCIAL, es decir, primero viene una instruccion, luego viene otra, y asi,
// El orden de ejecucion es siempre fijo linea por linea

// En cambio en la concurrencia, que es donde usamos hilos, esto ya no es mas asi, ya que en la concurrencia los hilos se ejecutan 'simultaneamente', en realidad 'turnandose' por lo
// que lo mas importante de la concurrencia es; EL PROCESO DE EJECUCION PUEDE SER DISTINTO PERO SIEMPRE EL RESULTADO FINAL ES EL MISMO

// Para entenderlo mejor veamos un ejemplo, para ello vamos a primero hacer una funcion comun y corriente

func imprimirCantidad(etiqueta string) {
	for cantidad := 1; cantidad <= 10; cantidad++ {
		fmt.Printf("Cantidad: %d de %s\n", cantidad, etiqueta)
	}
}

// Es bastante boluda la funcion, unicamente lo que le pasamos es un string por ejemplo imprimirCantidad("A")
// y lo que hace es imprimir 10 veces esa etiqueta con el numero de impresion

// Ahora vamos a usar la libreria SYNC luego explicaremos para que sirve cuando veamos la instruccion 'go'
// Aca vamos a crear una variable global que viene del paquete SYNC que es la que maneja todos los temas relacionados a la SINCRONIZACION
// Lo que hace es controlar que gorutinas se va a esperar a que se ejecuten
var wg sync.WaitGroup

// Ahora lo que vamos a hacer es a implementar la misma funcion imprimirCantidad() pero teniendo en cuenta que es una GORUTINA
// por tanto se llamara imprimirCantidad2()

func imprimirCantidad2(etiqueta string) {

	// Como sabemos que imprimirCantidad2() va a ser llamada como una GORUTINA y se ejecutara en un hilo aparte
	// lo que haremos a la variable global 'wg' es un DONE(), esto quiere decir que cuando se termine de ejecutar esta GORUTINA
	// se marcara un DONE como que esa GORUTINA ya finalizo
	defer wg.Done() // Le colocamos un 'defer' para quie siempre se ejecute inmediatamente antes de salir de la gorutina

	for cantidad := 1; cantidad <= 10; cantidad++ {
		// Ahora, por cada iteracion vamos a poner un sleep artificial, esto lo hacemos a fines pedagogicos
		// Para que se entienda bien la concurrencia como es que los hilos se van ejecutando 'simultaneamente'
		sleep := rand.Int63n(1000) // Lo que hace es tomar un numero random entre 0 y 1000
		// Ahora llamo a la funcion Sleep() del paquete TIME que lo que hace es detener la ejecucion de la GORUTINA
		time.Sleep(time.Duration(sleep) * time.Millisecond) // LO que hacemos es que ponga en pausa la ejecucion por lo que vale la variable 'sleep' en la unidad milisecond
		fmt.Printf("Cantidad: %d de %s\n", cantidad, etiqueta)
	}
}

// Ahora main...
func main() {
	fmt.Println("Iniciamos...")

	// Vamos a llamar la funcion con la etiqueta 'A'
	imprimirCantidad("A")

	// Ahora con la etiqueta "B"
	imprimirCantidad("B")

	// Esperamos a que finalice e imprimimos las sig. lineas:
	fmt.Println("Esperando que finalicen...")
	fmt.Println("Terminando el programa")

	fmt.Println("------------------------------------")

	// Si ejecutamos el programa vamos a ver que la ejecucion es SECUENCIAL, es decir, primero se ejecuta imprimirCantidad("A"), esperamos a que
	// termine de ejecutarse y luego leemos la instruccion imprimirCantidad("B"), termina de ejecutarse y luego se hacen los prints finales

	// Bueno todo esto ultimo es la ejecucion SECUENCIAL, es decir, lo que veniamos haciendo hasta ahora

	// Ahora que pasaria si quiero hacer que imprimirCantidad("A") e imprimirCantidad("B") se ejecuten 'simultaneamente'?
	// Bueno colocandole la palabra clave 'go' al principio

	// A esa variable 'wg' le tenemos que agregar la cantidad de gorutinas que queremos que espere, en este caso son 2
	// Ademas, si nos fijamos dentro de imprimirCantidad2() tenemos un Done(), esto quiere decir que se restara -1 a este Add() cuando la
	// gorutina termine de ejecurse en su respectivo hilo
	wg.Add(2)
	fmt.Println("Iniciamos las gorutinas en un hilo aparte...")

	go imprimirCantidad2("A") // Cuando le colocamos 'go' a esta instruccion quiere decir que esta instruccion en un hilo aparte se va a ejecutar mientras que el programa principal
	// sigue ejecutandose

	go imprimirCantidad2("B") // Lo mismo con esta linea de instruccion

	fmt.Println("Esperando que finalicen...")
	wg.Wait() // Aca ejecutamos la funcion wait() del paquete SYNC que lo que hace es que se esperara a que la cantidad de GORUTINAS que le pasamos
	// a 'wg' finalicen, que serian 2 gorutinas en total
	fmt.Println("Terminando el programa")

}

// Ahora vamos al archivo gorutinas2.md
