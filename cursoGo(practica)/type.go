package main

import "fmt"

// Ahora vamos a ver como declarar nuestros PROPIOS TIPOS DE DATOS en GO
// Para hacer esto tenemos que usar la palabra reservada TYPE

// Uno para declarar un nuevo tipo de dato ya toma un tipo existente, ejemplo:
type Dinero int

// Lo que estoy diciendo es que ahora Dinero es un tipo de dato INT!

// Entonces de ahora en mas 'Dinero' es un tipo de dato equivalente a INT

// Ahora lo que uno se puede preguntar ¿Entonces que sentido tiene si Dinero=INT? Entonces porque no use
// directamente el INT?

// Bueno, la magia esta en que yo puedo declarar METODOS que funcionan unicamente con esos tipos de datos, en este caso
// Dinero, por ejemplo:

func (d Dinero) String() string { // Esto no es una funcion, es un METODO, y se llama String() que es una palabra clave asociada al PrintLn(), lo que hace es recibir un tipo de dato Dinero 'd'
	// y lo que devuelve es un string
	return fmt.Sprintf("$%d", d) // Devuelve una string, recordar que Sprintf() sirve para formatear cadena

}

func main() {

	var sueldo Dinero
	sueldo = 25000

	// Entonces ahora cuando llamo a PrintLn() lo que va a hacer es llamar al metodo String() que formatea la cadena para que delante de todo haya un '$' para todo aquel tipo de dato
	// que sea Dinero como lo es 'sueldo'
	fmt.Println("El sueldo es de:", sueldo)

	// En cambio si a PrintfLn() le pasamos por ejemplo un INT entonces lo anterior no va a tener efecto
	// Es decir, el PrintLn() funciona comun y corriente:
	numero := 20
	fmt.Println("El numero es:", numero)

	// Otra cosa a tener en considerecion es que si bien Dinero parte de un INT no se puede sumar un INT con un tipo de dato Dinero:
	// sueldo += numero // Esto nos daria ERROR ya que estamos sumando 2 tipos de datos distintos si bien Dinero nacio de un INT no importa

	// Por lo tanto deberiamos hacer un casteo...
	sueldo += Dinero(numero)
	fmt.Println("El dinero actual es de:", sueldo)
}

// Ahora vamos al archivo estructuras.go
