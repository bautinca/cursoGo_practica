package main

import "fmt"

// Vamos a terminar de ver algunas cosas de estructuras

// Declaramos una struct Punto
type Punto struct {
	x, y int
}

// Otra cosa de las struct es que sus atributos no pueden ser del mismo tipo que la propia estructura que estamos definiendo
// osea:

type Punto3D struct {
	x, y, z int
	// Punto3D // Esto NO SE PUEDE HACER, lo que si se puede hacer es que, al igual que C sea un puntero a su mismo tipo, osea:
	*Punto3D
}

// Tambien podemos definir structs sin atributos
type OpPunto struct {
}

func main() {

	// Cual seria el valor de una struct declarada pero no definida? Logicamente es el valor 0 de c/u de sus atributos, osea:
	var p Punto
	fmt.Println("El valor de p es:", p) // Logicamente se imprimiria {0 0}

	// Ahora vamos a declarar y definir un Punto3D
	p2 := Punto3D{
		5,
		6,
		4,
		&Punto3D{
			6,
			4,
			6,
			nil,
		},
	}

	// Por tanto p2 uno de sus atributos es un puntero a otro Punto3D y luego ese otro Punto3D apunta a nil
	fmt.Println("El valor de p2 es:", p2)

	// Para acceder a ese otro Punto3D...
	fmt.Println("El valor p2 tiene otro Puntero3D que es:", *p2.Punto3D)

	// Tambien otra cosa es que las struct SON COMPARABLES SI SUS ATRIBUTOS SON COMPARABLES
	a := Punto{5, 6}
	b := Punto{7, 4}

	fmt.Println("a == b:", a == b) // Para que 2 estructuras sean iguales c/u de sus parametros deben ser iguales
}

// Ahora vamos a interfaces.go
