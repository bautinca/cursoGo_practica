package main

import "fmt"

// Ahora vamos a ver las famosas ESTRUCTURAS!, en C son STRUCTS (si bien C no esta
// especialmente diseñado para POO), y en Python tenemos las CLASS

// Aca las estructuras funcionan casi identico a en C, por ejemplo hagamos la STRUCT de una Persona:

type Persona struct {
	Nombre string // La struct Persona
	Edad   int
}

func main() {

	// Vamos a declarar una variable de tipo Persona
	var p Persona
	// Ahora 'p' es un TDA Persona, por lo que le podemos definir los atributos igual que en C:
	p.Nombre = "Alejandro"
	p.Edad = 29

	fmt.Println("Estructura p de tipo Persona", p) // Se imprime como {Alejandro 29} la 'p'
	fmt.Println("El nombre de p es:", p.Nombre)
	fmt.Println("La edad de p es:", p.Edad)

	// Tambien podemos declarar y definir a un TDA Persona todo en la misma linea asi:
	p2 := Persona{Nombre: "Rafael", Edad: 25}

	fmt.Println("Estructura p2 de tipo Persona", p2)

	// Tambien hay una tercer manera de definir y declarar las TDA que es colocando directamente el valor
	// de los atributos (si ya nos sabemos el orden de los atributos en el struct)
	p3 := Persona{"Miguel", 18}
	fmt.Println("Estructura p3 de tipo Persona", p3)
}

// Ahora vamos a estructuras2.go
