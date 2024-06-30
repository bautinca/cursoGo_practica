package main

import "fmt"

// Creamos un struct Persona

type Persona struct {
	Nombre   string
	Apellido string
}

// Ahora creamos otro struct donde uno de sus atributos es justamente un TDA Persona!
type Estudiante struct {
	Persona
	Carrera string
}

func main() {

	// Ahora vamos a declarar y definir una variable
	alejandro := Estudiante{
		Persona{
			Nombre:   "Alejandro",
			Apellido: "Arnaud",
		},
		"Informatica",
	}

	// Ahora imprimimos a 'alejandro'
	fmt.Println(alejandro)

	// Ahora bien, la novedad es que la estrucutra Estudiante tiene DIRECTO ACCESO a todos los atributos de la estructura Persona,
	// Es decir:

	fmt.Println("El nombre del estudiante es:", alejandro.Persona.Nombre)
	fmt.Println("El nombre del estudiante es:", alejandro.Nombre)

	// Es decir, es exactamente lo mismo podemos hacer, en ambos casos se imprime exactamente lo mismo
}

// Ahora vamos al archivo estructuras3.go
