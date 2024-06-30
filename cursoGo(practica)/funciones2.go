package main

import "fmt"

// Nosotros hasta ahora vimos funciones comunes y corrientes, que se las pueden ver como 'paquetes'
// Ahora el tema es que no puedo definir una funcion dentro de otra funcion, por ejemplo dentro del main
// no puedo crear una funcion

// Pero si quiero definir una funcion dentro de otra entonces como hago? Para ello existe el concepto de CLOSURE
// Bueno, si recordamos el 'func' es un TIPO DE DATO!, entonces las funciones son valores!

// Por ejemplo supongamos que tenemos la sig. funcion
func print(cadena string) {
	fmt.Println(cadena)
}

// Y tengo la sig. funcion que es exactamente igual a la anterior pero con otro nombre
func print2(cadena string) {
	fmt.Println(cadena)
}

// Y esta otra funcion
func print3(cadena1, cadena2 string) {
	fmt.Println(cadena1 + cadena2)
}

// Ahora como las funciones SON VALORES entonces una funcion puede recibir otra funcion!!!
func print4(f_print func(string)) { // Estamos diciendo que la funcion print4() recibe una funcion llamada f_print que recibe un string
	f_print("Hola queridos mios")
}

// Ahora vamos a ver una funcion particular, que como las funciones son valores entoncespodemos hacer que una funcion devuelva otra funcion, es decir:
func incrementar() func() int { // Aca estamos diciendo que incrementar() es una funcion que no recibe parametros y devuelve una funcion que no recibe tampoco parametros y devuelve un INT
	i := 0
	return func() (r int) { // Estoy diciendo que la funcion incrementar() retorne una funcion que retorna 'r' que es un int
		r = i
		i++
		return // se retorna r
	}

}

// Es decir, lo que hice con esta funcion incrementar() es exactamente definir una funcion dentro de otra,
// es decir, incrementar() no recibe parametros y devuelveuna funcion que tampoco recibe parametros y devuelve un int
// Luego dentro de incrementar() logicamente retorno una funcion que retorna un int definido en 'r' entonces
// lo que hago debajo del primer return es exactamente definir la funcion que retorna! y cuando defino esta funcion logicamente en el segundo return retorna a 'r'!

func main() {

	// Dentro de main vamos a declarar y definir una variable que sera un string
	cadena := "Hola como estan"

	// Ahora bien, como FUNC es un tipo de dato eso quiere decir que LAS FUNCIONES SON VALORES!
	imprimir := print

	// Por lo tanto ahora 'imprimir' es de tipo de dato FUNC, es decir, es una funcion! y hace lo mismo que print()
	imprimir(cadena)

	// Ahora, para definir una funcion dentro de otra, por ejemplo aca que estamos dentro de main() se hace asi:
	// Ahora diremos que imprimir2 sea una funcion que la declararemos y la definiremos aca dentro:
	imprimir2 := func() {
		fmt.Println(cadena) // Podemos usar el parametro 'cadena' ya que esta definido dentro del main(), y como esta funcion esta dentro del main() entonces puede acceder
		// a dicha variable 'cadena'
	}

	// Y listo! ahora imprimir2() es claramente una funcion! que lo unico que hace es imprimir 'cadena':
	imprimir2()

	// Bien acabamos de hacer un CLOSURE!

	// Ahora vamos a redefinir 'imprimir'
	imprimir = print2 // Ahora 'imprimir' hace lo que hace print2()
	// Como 'imprimir' hace lo que hace print2 entonces es una funcion imprimir() y recibe la cadena que va a imprimir
	imprimir("Esto es una locura")

	// Es decir, le cambiamos el valor a 'imprimir'
	// Ahora le cambiaremos nuevamente el valor a 'imprimir' por print3()

	// (descomentar)
	//imprimir = print3 // ATENCION ACA ¿Se puede igualar 'imprimir' a la funcion print3? NOOOOO
	// ¿Pero porque si antes 'imprimir' equivalia a print() y despues a print2(), porque con print3() no me deja?
	// Bueno esto es por un tema de FIRMAS, es decir, si nos fijamos print() y print2() TIENEN EXACTAMENTE LA MISMA FIRMA, es decir,
	// reciben 1 parametro string y devuelven lo mismo, osea nada
	// En cambio la firma de print3() recibe 2 STRINGS, por ende su firma es distinta y es por eso que 'imprimir' no puede equivaler a print3()
	// Porque al momento que definimos a imprimir:=print es en ese momento que de ahora en mas 'imprimir' solo puede equivaler a funciones que reciben 1 parametro y no devuelven nada

	// NOTA: Para ver las firmas de las funciones podemos usar %T!!, osea:
	fmt.Printf("La firma de print es: %T\n", print)
	fmt.Printf("La firma de print2 es: %T\n", print2)
	fmt.Printf("La firma de print3 es: %T\n", print3)

	// Efectivamente al ver la STDOUT la firma de print3 es distinta a la de print y print2

	// Ahora vamos a ver el print4() esta es una funcion que recibe como parametro otra funcion que esta otra funcion recibe un string,
	// por lo tanto se podria hacer:
	print4(print) // Le podemos pasar la funcion print() porque claramente print() recibe 1 solo string como parametro entonces
	// lo que hara el print4() es en su definicion que hicimos, trabajar con la funcion print()

	// Ahora bien, como las FUNC son un tipo de dato entonces podemos declararlas como una variable comun y corriente, osea:
	var fb func()
	// En este caso fb es una funcion VACIA, osea NIL ya que no esta definida
	fmt.Println(fb) // Claramente imprime NIL

	// Ahora bien, si las funciones son valores estos NO SON COMPARABLES, unicamente son comparables con NIL ya que o son NIL o no son NIL

	// Ahora bien, vamos a ver la funcion incrementar() que esta funcion no recibe nada y devuelve una funcion que esta funcion que devuelve devuelve un int, es decir, definimos una funcion
	// dentro de la funcion incrementar(). Entonces lo que haria incrementar() es como una especie de contador que cada vez que la llamamos se incrementa +1 en de la funcion que definimos dentro de incrementar()

	// Antes que nada incrementar() sabemos que es un valor, por lo tanto se lo asignamos a una variable
	contador := incrementar()

	// Ahora corremos el contador
	fmt.Println(contador()) // 0
	fmt.Println(contador()) // 1
	fmt.Println(contador()) // 2
	fmt.Println(contador()) // 3
	fmt.Println(contador()) // 4
	fmt.Println(contador()) // 5

	// Y listo!
}

// Ahora vamos a funciones3.go
