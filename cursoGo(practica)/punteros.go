package main

import "fmt"

// Esta funcion leerla cuando llegamos a la parte de punteros en funciones
func incrementar(n int) {
	n++
	fmt.Println("El valor incementando es:", n)
}

// Esta funcion leerla cuando llegamos a la parte de punteros en funciones
func incrementar2(n *int) {
	*n++
	fmt.Println("El valor incrementado es:", *n)
}

func main() {

	// Ahora vamos a hablar de los famosos PUNTEROS! Donde el concepto es tal cual identico que en C,
	// es decir, un puntero apunta a la direccion de memoria de una variable, y solo eso!

	// Declaramos una variable
	a := 25

	// Ahora vamos a imprimir el valor de a
	fmt.Println("El valor de 'a' es:", a)

	// Por otro lado vamos a imprimir su direccion, que para sacar su direccion, al igual que en C, tenemos que colocarle el &
	fmt.Println("La direccion de la variable 'a' es:", &a)

	// Ahora vamos a declarar un puntero, vamos a decir que b->a, osea la variable 'b' apuntaria a la direccion de la variable 'a'
	b := &a // Ahora 'b' es un puntero!

	// Por lo tanto es logico que el valor del puntero 'b' es justamente la direccion de 'a'
	fmt.Println("El valor de 'b' es:", b)

	// Ahora al igual que en C, si al puntero le colocamos un * delante, quiere decir que queremos acceder AL VALOR hacia donde apunta el puntero
	// Es decir que *b seria equivalente al valor de 'a'
	fmt.Println("El puntero 'b' apunta al valor:", *b)

	// Ahora bien, es importante decir, que tambien al igual que C, 'b' es un PUNTERO, no un entero, eso quiere decir que su tipo de
	// dato es INT* y no INT, por lo tanto hacer b=10 esta MAL

	// Si quiero asignarle un nuevo valor a 'a' a traves de su puntero lo que puedo hacer es...
	*b = 11 // Con esto lo que hacemos es cambiarle indirectamente el valor a la variable 'a'
	fmt.Println("El valor modificado de 'a' es:", a)

	// Ahora logicamente si por ejemplo modificamos el valor de 'a' de modo directo entonces eso tambien se va a modificar en el puntero 'b' hacia que valor apunta
	a++ // Incrementamos +1 a 'a'
	fmt.Println("El puntero 'b' apunta hacia:", a)

	// Ahora una cosa a saber es que LOS PUNTEROS son comparables, es decir, por ejemplo 2 punteros son equivalentes cuando apuntan hacia la misma direccion, por ejemplo
	// supongamos que ahora tengo un nuevo puntero que apunta hacia la direccion de 'a'
	c := &a

	// Entonces logicamente los punteros 'b' y 'c' deberian ser equivalentes
	if b == c {
		fmt.Println("Los punteros 'b' y 'c' son equivalentes")
	}

	// Ahora algo nuevo que tenemos en GO es que con new() declaramos un puntero donde dentro le pasamos de que tipo de dato
	d := new(int) // En este caso 'd' esta declarado como un INT*

	// A que direccion apuntaria un puntero solo declarado pero no definido?
	fmt.Println("La direccion a la que apunta 'd' recien creado es:", d)
	// Y a que valor apunta?
	fmt.Println("El valor al que apunta el puntero 'd' es: ", *d)

	// Si nos fijamos logicamente apuntaria a un valor 0 cuya direccion ya esta definida
	// Ahora lo que voy a hacer es igualar el puntero 'd' con el de 'b', por lo tanto 'd' apuntaria hacia la direccion de 'a'
	d = b
	// Ahora consultamos hacia que valor apunta 'd', logicamente tiene que ser el de la variable 'a', osea 12
	fmt.Println("El valor hacia el que apunta el puntero 'd' es ahora:", *d)

	// Ahora vamos a ver como pasar punteros a funciones, los punteros vimos gracias a C que son muy buenos para cambiar el estado de las cosas
	// Por ejemplo, supongamos que ahora tenemos la sig. variable
	numero := 0

	// Ahora supongamos que quiero crear una funcion que tome a 'numero' y le vaya incrementando +1
	// Se puede hacer esto? Por ejemplo:
	incrementar(numero)
	incrementar(numero)
	incrementar(numero)

	// Porque en los 3 casos al incrementarle +1 a 'numero' se imprime 1? Porque Lo que sucede es que a cada incrementar
	// le entra el numero=0, es decir, no vamos subiendo en una unidad a la variable 'numero' cada vez que llamamos a incrementar()

	// Entonces para solucionar esto esta el PUNTERO! que ahora haremos un nuevo incrementar2() que este si recibe el puntero de 'numero'
	incrementar2(&numero)
	incrementar2(&numero)
	incrementar2(&numero)

	// Ahora si funciona!
}

// Ahora vamos al archivo type.go
