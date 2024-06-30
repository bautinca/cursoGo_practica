package main

import (
	"errors"
	"fmt"
)

// Una de las cosas que tiene GO es que no lanza ERRORES, es decir, nos da mas libertad para que NOSOTROS manejemos
// los errores. Para ello tenemos un paquete que se llama 'errors'

// package errors

// Algo interesante que tiene el paquete errors es que tenemos la funcion new() donde se nos permite crear nuestro propio error
// De hecho, como ya hemos visto, el ERROR en si es un tipo de dato

// Lo que vamos a hacer es crear una funcion que recibe un string y devuelve un error, logicamente
// no tiene sentido que una funcion devuelva un error, pero la hacemos asi a fines pedagogicos

func baneado(usuario string) (err error) {
	// Por default ban equivale a false
	ban := false

	switch usuario {
	case "miguel":
		ban = true // Recordar que cuando se termina de ejecutar un case automaticamente salimos del switch, ni el default se ejecuta

	case "carlos":
		ban = false

	case "juan":
		return fmt.Errorf("El usuario no es valido") // Retornamos en este case logicamente un error con Errorf()

	case "pedro":
		return fmt.Errorf("Usuario en proceso de registro") // Devolvemos un error

	default:
		return fmt.Errorf("Algo paso...") // Devolvemos un error
	}

	if !ban { // Si ban=false
		fmt.Printf("El usuario %s no esta baneado\n", usuario)
	} else {
		fmt.Printf("EL usuario %s esta baneado\n", usuario)
	}

	return nil // Retornamos logicamente un error, donde un error tambien puede equivaler a nil que es cuando no tenemos error

}

// Aca implementamos la funcion checkError() para tratar los errores
func checkError(err error) { // Recibe logicamente un error
	// Y aca colocamos el codigo que se suele repetir
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// Ahora, nosotros habiamos dicho que podemos crear errores con new() ¿Como los creamos? Para ello lo que vamos a hacer es crearnos 3 variables
// de tipo ERR

var ErrorUsuarioNoValido error = errors.New("El usuario no es valido")
var ErrorUsuarioEnProceso error = errors.New("Usuario en proceso de registro")
var ErrorPorDefecto error = errors.New("Algo paso...")

// Y lo que haremos es redefinir baneado() pero con estos errores que creamos, baneado2() seria

func baneado2(usuario string) (err error) {
	// Por default ban equivale a false
	ban := false

	switch usuario {
	case "miguel":
		ban = true // Recordar que cuando se termina de ejecutar un case automaticamente salimos del switch, ni el default se ejecuta

	case "carlos":
		ban = false

	case "juan":
		return ErrorUsuarioNoValido

	case "pedro":
		return ErrorUsuarioEnProceso

	default:
		return ErrorPorDefecto
	}

	if !ban { // Si ban=false
		fmt.Printf("El usuario %s no esta baneado\n", usuario)
	} else {
		fmt.Printf("EL usuario %s esta baneado\n", usuario)
	}

	return nil // Retornamos logicamente un error, donde un error tambien puede equivaler a nil que es cuando no tenemos error

}

// Ahora la funcion main()
func main() {
	err := baneado("miguel") // En este caso 'miguel' si esta baneado
	if err != nil {          // Si hay error..., logicamente baneado("miguel") retorna un nil
		fmt.Println("Error:", err)
	}

	// Ahora lo que vamos a hacer es llamar por c/u de los nombres
	err = baneado("carlos") // En este caso carlos no esta baneado por tanto la funcion retorna nil
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Ahora con pedro
	err = baneado("pedro") // EN este caso retornaria un error donde se imprime Uusario en proceso de registro
	if err != nil {        // Logicamente vamos a entrar a este if para imprimir el error que retorna
		fmt.Println("Error:", err)
	}

	// Finalmente un nombre cualquiera
	err = baneado("pelelo") // Esto retorna un error en el que algo paso, por tanto entramos al if
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Ahora lo que notamos es que siempre colocamos el codigo siguiente:
	/*
		if err != nil {
			fmt.Println("Error:", err)
		}
	*/

	// Entonces para evitar repetir esas lineas de codigo cientos de veces lo que podemos hacer es ENCAPSULARLA EN UNA FUNCION QUE CHEQUEA EL ERROR!, para
	// ello esta la funcion que implementamos arriba del main() checkError()
	fmt.Println("----------------------------------------")

	// Ahora tan solo lo que tenemos que hacer es siempre inmediatamente abajo de cada llamada a baneado() verificamos si dio
	// error con el checkError()

	err = baneado("miguel")
	checkError(err)

	err = baneado("carlos")
	checkError(err)

	err = baneado("pedro")
	checkError(err)

	err = baneado("pelelo")
	checkError(err)

	fmt.Println("------------------------------------------")

	// Ahora vamos a hacer lo mismo pero usando baneado2() donde implementamos nuestros propios errores, pero deberia funcionar tal cual igual
	err = baneado2("miguel")
	checkError(err)

	err = baneado2("carlos")
	checkError(err)

	err = baneado2("pedro")
	checkError(err)

	err = baneado2("pelelo")
	checkError(err)
}

// Ahora vamos a ir a gorutinas.go
