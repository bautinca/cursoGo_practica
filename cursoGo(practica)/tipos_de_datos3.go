package main

import (
	"fmt"
)

func main() {

	// Ahora vamos a ver todos lo que tenga que ver con LOS DATOS DE TIPO FLOTANTE

	// El primero es el float32 donde ocupara 32 bits = 4 bytes, por tanto tiene 6-9 digitos decimales
	// de precision
	var f32 float32 = 5.42373285723
	fmt.Println(f32) // Logicamente se va a imprimir con 6-9 digitos decimales de precision, osea no todos, osea 5.4237328 en este caso imprime

	// Ahora el float64 que la logica es la misma, pero ahora tenemos 64 bits a disposicion, por lo tanto la precision va a ser mucho mayor
	var f64 float64 = 5.42373285723
	fmt.Println(f64) // Aca logicamente se van a imprimir mas decimales, de hecho todos en este caso

	// Una novedad que tenemos en Go es que tenemos tipos de datos para numeros complejos
	var c64 complex64   // Numero complejo de 64 bits
	var c128 complex128 // Numero complejo de 128 bits

	// Ahora vamos a ver como se imprimen los numeros complejos, obviamente si no los definimos por default quedan en 0
	// Por lo tanto, siguiendo la logica ambos (c64 y c128) se tendrian que imprimir como 0+0i
	fmt.Println(c64, c128)

	// Logicamente con los numeros flotantes tenemos los mismos OPERADORES que con los numeros enteros
	// es decir, los podemos sumar, restar, comparar, incrementar, etc.

	// Como podemos asignarle valor a los numeros complejos? Asi:
	c64 = 4 + 3i
	fmt.Println(c64)

	// Y la otra manera:
	c128 = complex(10, 12)
	fmt.Println(c128)

	// Ademas con real() sacamos la parte real e imag() la parte imaginaria
	fmt.Println(real(c64))
	fmt.Println(imag(c64))
}

// Ahora vamos al archivo tipos_de_datos4.go
