package main

import "fmt"

func main() {

	// En GO logicamente tenemos los mismos operadores que en todos los lenguajes de programacion, o la
	// gran mayoria, como en C y Python, osea:

	fmt.Println(5 + 3)
	fmt.Println(5 - 3)
	fmt.Println(5 * 3)
	fmt.Println(5 / 3) // Como es una operacion entre enteros entonces logicamente el resultado me va a dar otro entero, Go como
	// en cualquier otro lenguaje quitaria el decimal
	fmt.Println(5 % 3)

	// Ahora los operadores de COMPARACION, que logicamente son los mismos que la gran mayoria de lenguaje
	numero1 := 5
	numero2 := 3

	fmt.Println(numero1 == numero2)
	fmt.Println(numero1 != numero2)
	fmt.Println(numero1 < numero2)
	fmt.Println(numero1 > numero2)
	fmt.Println(numero1 <= numero2)
	fmt.Println(numero1 >= numero2)

	// En GO como en la gran mayoria de lenguajes, tenemos los operadores
	// += o -= !! Es decir, en lugar de hacer a = a + b podemos hacer a += b
	numero1 += numero2
	fmt.Println(numero1) // Se imprimiria 8

	// Tambien logicamente tenemos los OPERADORES LOGICOS que son tambien los mismos que en cualquier otro lenguaje, pero
	// aca podemos usar los operadores logicos de Python o si queremos de C
	fmt.Println(true || false) // Operador OR
	fmt.Println(true && false) // Operador AND
	fmt.Println(!true)         // Operador NOT

	// En GO tenemos como en C, los operadores de INCREMENTO o DECREMENTO en +1 o -1
	numero2++
	fmt.Println(numero2)
	numero2++
	fmt.Println(numero2)
	numero2--
	fmt.Println(numero2)
}

// Ahora vamos a ir al archivo tipos_de_datos3.go
