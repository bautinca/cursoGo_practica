package main

import "fmt"

// Hasta ahora veniamos trabajando siempre con la funcion principal MAIN, pero ahora vamos a ver como
// crear funciones en general

// Es bastante igual al resto de los lenguajes de programacion, aca para declarar una funcion usamos la palabra FUNC y en los parametros
// que recibe tenemos que indicar de que tipo son los parametros que recibe, es decir, igual que C en ese sentido

func imprimiNombre(nombre string) {
	fmt.Println("Fuera del main")
	fmt.Println("El nombre es:", nombre)
}

// Y listo! Nuestra primera funcion, aunque si nos fijamos no retorna nada

// Ahora hagamos otro tipo de funcion que esta si va a retornar algo, precisamente un INT, entonces al igual que en C, tenemos que especificar que tipo de dato
// va a retornar
func suma(n1 int, n2 int) int { // Esto quiere decir que retorna un INT
	return n1 + n2
}

// Ahora, una novedad que tiene GO es que podemos especificar que variable queremos que retorne y su tipo de dato, por ejemplo:
func resta(n1, n2 int) (r int) { // Varias cosas a aclarar, si todos los parametros que recibe la funcion son del mismo tipo, entonces se puede declarar el tipo 1 vez, como n1 y n2 son de tipo
	// int entonces se declara que son de tipo int 1 vez podemos hacer
	// Por otro lado estamos declarando una variable nueva dentro de la funcion que es 'r' que es de tipo INT, y que ademas ese mismo 'r' sera el resultado que retorne la funcion

	r = n1 - n2
	return // Automaticamente va a retornar 'r'
}

// Ahora, tambien podemos hacer que las funciones nos den 2 resultados! Osea tal cual como con los maps que al seleccionar una clave nos
// daba un segundo resultado (BOOL) donde nos decia si dicha clave pertenecia al map o no

func division(n1, n2 float64) (r float64, err error) {
	if n2 == 0 {
		err = fmt.Errorf("No se puede dividir por cero")
		return // Este return nos devolveria r=0 y err=fmt.Errorf(...)
	}

	r = n1 / n2
	return // Este resultado nos devolveria r=n1/n2 y err=nil
}

// Ahora vamos a ver una funcion particular que se llaman VARIADIC FUNCTIONS, estas son funciones
// que pueden recibir 'x' argumentos, osea la cantidad de argumentos que recibe es una variable, por ejemplo yo puedo hacer suma(3), o suma(5,4), o suma(5,4,2,5,2)
// o suma(5,321,563,2,326,23,5,235,6,243,32) todo en la misma funcion sin especificar la cantidad de argumentos que recibe

// En GO para declarar este tipo de funciones se usa la sintaxis '...', es decir, aquel argumentos que queremos que sea una variable la indicamos con ..., osea asi:
func sumarNumeros(mensaje string, numeros ...int) int {
	// Aca el parametro VARIADICO 'numeros' se comporta como un SLICE, osea imaginar que es un SLICE, por lo tanto lo podemos recorrer, seria un slice de INTs
	fmt.Println(mensaje)
	suma := 0
	// Como 'numeros' se comporta como un slice lo podemos recorrer con RANGE
	for _, numero := range numeros { // Omitimos el indice ya que no lo necesitamos
		suma += numero
	}
	return suma
}

// Bueno listo, ya hicimos nuestras 3 funciones auxiliares, ahora en la funcion principal main() vamos a probarlas
func main() {
	imprimiNombre("Pepe")
	fmt.Printf("El resultado de 5 + 3 es: %d\n", suma(5, 3))
	fmt.Printf("El resultado de 5 - 3 es: %d\n", resta(5, 3))

	// Y ahora la division que devuelve 2 resultado, el resultado de la division y si tenemos error o no
	resultadoDivision, err := division(5.0, 0)
	if err == nil {
		fmt.Println("El resultado de 5/2.5 es:", resultadoDivision)
	} else {
		fmt.Println("Error", err)
	}

	// Aca vamos a evaluar la VARIADIC FUNCTION
	fmt.Printf("El resultado de 5+2+10+42+53+2 es: %d\n", sumarNumeros("calculando...", 5, 2, 10, 42, 53, 2))

}

// Ahora vamos a ir al archivo funciones2.go
