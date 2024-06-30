package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// Ahora del tipo de dato que vamos a hablar son de los STRINGS
	// El tipo de dato STRING es basicamente un arreglo de CHARs si lo queremos ver desde el punto
	// de vista de C
	var cadena string = "Hola como estas"
	fmt.Println(cadena)
	fmt.Println(unsafe.Sizeof(cadena)) // En este caso la cadena pesa 16 bytes, obviamente mientras mas larga sea la cadena mas va a ocupar

	// En GO, a diferencia de en C, no existe el tipo de dato para un unico caracter, osea un CHAR

	// Tambien al ser un arreglo al fin y al cabo, le podemos sacar su largo:
	fmt.Println(len(cadena))

	// Y tambien podemos seleccionar el dato de cualquiera de sus indices o grupos de indices:
	fmt.Println(cadena[2])   // Porque me imprime el 108? Porque justamente el 108 representa el caracter 'l' que en Unicode en decimal es el 108
	fmt.Println(cadena[1:6]) // Toma entre el indice 1 (inclusive) y el indice 6, recordemos que esto de los slices solo se puede hacer en Python!, el concepto
	// de SLICES en GO es tal cual identico que en el de Python, podriamos hacer tambien cadena[:4] o cadena[::1] por ejemplo, es decir, todas las opciones de slices

	// Ahora bien, algo a destacar de los tipos de dato STRING es que son INMUTABLES, es decir, no se pueden modificar
	// es decir, no podemos hacer cadena[1] = "x" por ejemplo

	// Tambien otra de las cosas es que las cadenas las podemos concatenar
	cadena = cadena + " pepito"
	fmt.Println(cadena)

	// Tambien se podrian concatenar asi las cadenas:
	cadena += " y jose"
	fmt.Println(cadena)

	// Ahora, una novedad que tiene GO, es que, obviamente como vemos hasta ahora todas las cadenas se definen entre "" (comillas dobles)
	// PEEEROO, ahora lo nuevo es que en GO podemos usar "`", esto es para que una cadena tome literalmente esa forma, es decir:

	var cadena2 string = `

Esto es una cadena
	mas xx
		compleja 294239
	se va #$ a imprimir      todo esto
tal cual   --}~ coo lo escribo
`

	fmt.Println(cadena2)

	// Tambien GO soporta los caracteres especiales como el salto de linea
	var cadena3 string = "Hola como\nEstas genio\tsisi" // el \t es una tabulacion
	fmt.Println(cadena3)

	// Ahora la gran duda es ¿Como concatenar strings con variables? Para ello tenemos que usar
	// el paquete STRCONV y de ella usar la funcion ITOA() donde justamente lo que hace es convertir una variable entero a un string y asi poder concatenarla con mas cadenas
	var entero int = 4
	cadena4 := "Esta es la cadena numero " + strconv.Itoa(entero) + ", no es un numero tan alto"
	fmt.Println(cadena4)

	// Tambien otra manera de pasar los enteros a cadenas es con Sprintf() del paquete fmt
	// Aca lo bueno es que si ya sabemos manejar C es una boludez saber como funciona, veamos:

	fmt.Println(fmt.Sprintf("Tengo %d copas", entero))

	// Como vemos Sprintf() es lo mismo que el printf() en C!

}

// Ahora vamos al archivo for.go
