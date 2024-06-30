package main

import (
	"fmt"
	"strings"
)

// Vamos a declarar aca una funcion anonima, para ello tenemos que crear una funcion comun y corriente
// que dentro tenga una funcion anonima (sin nombre), entonces cuando se ejecute esta funcion si o si va a ejecutar la anonima
func ejecutarAnonima(fn func(int, int) int, a, b int) { // Para entender, esta funcion recibe una funcion 'fn' que recibe 2 parametros INTs y devuelve un parametro INT y ademas recibe 2 INTs a y b
	resultado := fn(a, b)
	fmt.Println("El resultado es:", resultado)
}

func main() {

	// Ahora vamos a ver lo que son las funciones ANONIMAS
	// Nosotros inconcientemente ya vimos funciones anonimas, ¿Donde? Pues cuando definimos el incrementar() en funciones2.go

	// Es decir, si nos fijamos el incrementar() retorna UNA FUNCION ANONIMA ya que no tiene nombre!

	// Lo que vamos a hacer es usar el paquete STRINGS la funcion Map() que este si es el map() que vimos
	// en python, donde a cada elemento de una cadena le aplica una funcion especifica

	// Entonces, lo que vamos a decir es que esta funcion que recibe Map() es justamente anonima! (es como la funcion LAMBDA)

	// Primero definimos un string random de solo numeros
	cadena := "22847218432874299878912"

	// Ahora, vamos a actualizar 'cadena' por su valor luego de que pase por la funcion Map()
	cadena = strings.Map(func(r rune) rune {
		return r + 1
	}, cadena)

	// Que estamos diciendo? Bueno la funcion LAMBDA del Map() es una funcion que recibe el parametro
	// 'r' que es un tipo de dato rune = uint32 y devuelve otro tipo de dato rune

	// Por tanto lo que hacemos dentro del Map() es logicamente declarar la funcion lambda y definirla con 'r+1', por lo tanto el Map() recorre
	// cada elemento y le va sumando +1 y eso es lo que retornara

	// Logicamente como segundo parametro del Map() es la string que queremos recorrer
	fmt.Println(cadena) // Printea 339583295439853::989:23

	// Mas facil para entender lo de las funciones anonimas son al fin y al cabo funciones que NO TIENEN NOMBRE
	// es decir no se declaran por separado como una funcion comun y corriente, ya que estas funciones anonimas suelen
	// estar dentro de otras funciones que estas otras funciones (que no son anonimas) cuando se las use automaticamente van a ejecutar la funcion anonima

	// Por ejemplo dentro de este main() (que es una funcion al fin y al cabo) vamos a declarar una funcion anonima
	// Esta funcion anonima se la asignamos a la variable 'suma' ya que como sabemos, toda funcion es un tipo de dato
	suma := func(a, b int) int {
		return a + b
	}

	// Y listo!
	fmt.Println(suma(5, 3))

	// Ahora ejecutamos la otra funcion ejecutarAnonima() que toma una funcion anonima y la ejecuta
	ejecutarAnonima(suma, 10, 3)

}

// Ahora vamos al archivo defer.go
