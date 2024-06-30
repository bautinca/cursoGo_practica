/*
Bueno en este archivo vamos a hablar todo sobre las VARIABLES en GO, aca la VARIABLE
se la trata como en cualquier otro lenguaje de programacion, osea una variable es una variable y punto, con su nombre de identificacion, su respectivo valor
y obviamente esta variable se va a almacenar en la memoria con sus respectivas direcciones

Primero lo que vamos a hacer es declarar el paquete main
*/
package main

// Importamos ahora el paquete fmt para poder hacer impresiones
import "fmt"

// Ahora la funcion main()
func main() {
	// Para declarar una variable se usa la palabra 'var' y el nombre de la variable y luego el TIPO de dato de dicha variable
	var numero int

	// Ya declarada la variable 'numero' ahora si la podemos definir, osea es bastante similar a C
	numero = 25

	//Ahora imprimimos el valor 'numero'
	fmt.Println(numero)

	// Ahora lo que vamos a hacer es modificarle el valor a la variable 'numero' por otro valor e imprimimos nuevamente
	numero = 40
	fmt.Println(numero)

	// Si ejecutamos el programa vamos a ver que se imprime el 25 y luego el 40 y es logico
	// que suceda asi, ya que a la variable 'numero' le cambie de valor en la misma direccion de la memoria, y siempre el nombre 'numero' va a hacer alusion
	// a la misma direccion de memoria. Es decir, lo que queremos decir es que todo sucede en LA MISMA DIRECCION DE MEMORIA de que se modifica el valor, no es que se crea una direccion
	// para el valor 25 y otra direccion para el valor 40

	// Ahora si recordamos en INTRODUCCION.md habiamos dicho que el lenguaje GO tiene DUCKTYPING, esto es que no hace falta
	// declararle el tipo de dato a la variable ya que si por ejemplo x="pepe" ya se sobreentiende que x es un tipo de dato arreglo de CHARs
	// Y para hacer el ducktyping se usa la asignacion ':='
	nombre := "Alejandro"
	// Y lo imprimimos
	fmt.Println(nombre)

	// ¿Como declarariamos una variable que en C seria un array de chars? Osea una string al fin y al cabo? Asi:
	var arreglo string = "Pepe" // Si nos fijamos declaramos y definimos la variable todo en la misma linea, esto al igual que en C tambien se puede hacer
	// Y lo imprimimos...
	fmt.Println(arreglo)

	// Tambien lo que hay que entender del operador ':=' es que automaticamente
	// te asigna el tipo de dato, es decir, por ejemplo:
	x := 100
	// Entonces de ahora en mas siempre 'x' es de tipo de dato entero, no es que despues puede ser un char
	// Ahora hacemos un print...
	fmt.Println(x)

	// Ahora tambien le podemos modificar el valor a 'x' pero obviamente por un valor del mismo tipo de dato
	x = 110
	fmt.Println(x)

	// Entonces si o si 'x' es una variable de tipo entero, no le podemos cambiar el tipo de dato
	// es decir, no podriamos hacer x = "hola", esto estaria mal

	// Al igual que en Python, en GO podemos asignar en una misma linea varias variables
	nombre, numero = "Pepe", 10
	// Ahora imprimimos ambas variables, ademas si notamos directamente pusimos '=' sin declarar ninguna variable
	// esto esta permitido porque las variables 'nombre' y 'numero' ya fueron anteriormente declaradas
	fmt.Println(nombre, numero)

	// Ahora tambien lo que podemos hacer es un switch, por ejemplo:
	// Primero definimos una variable nueva nombre2:
	nombre2 := "Juan"

	// Ahora hacemos un switch entre las variables 'nombre' y 'nombre2'
	nombre, nombre2 = nombre2, nombre

	// Ahora hacemos un print para ver si se switchearon los valores correctamente
	fmt.Println(nombre, nombre2)

	// Efectivamente se switchearon los valores de 'nombre' y 'nombre2' y en ese sentido esto es igual a Python, y recordemos que para hacer esto en C
	// era mucho mas quilombo ya que teniamos que usar punteros

	// Tambien otra de las cosas que tiene GO es que TODAS LAS VARIABLES DECLARADAS DEBEN SER UTILIZADAS, porque caso contrario GO a la hora de compilar
	// nos lanzaria error

	// Ahora una de las cosas que tiene GO, es que si decidimos usar una variable que es un entero
	// que la declaramos pero no la definimos, entonces por default GO le asigna el valor de 0 a esa variable
	var entero int
	fmt.Println(entero)

	// Efectivamente el valor de 'entero' es 0
	// Por otro lado a las cadenas de caracteres declaradas pero sin definir les da una cadena vacia por default
	var cadenita string
	fmt.Println(cadenita)

	// A los booleanos por defaut les pone false si no estan definidos
	var booleano bool
	fmt.Println(booleano)

	// Y a los tipo de dato flotante por default les asigna 0.0 logicamente si no las definimos
	var flotante float64
	fmt.Println(flotante)

	// Otra de las cosas que no explicamos es que por ejemplo podemos hacer lo sig:
	var numero3 = 75
	fmt.Println(numero3)
}

// Ahora vamos al archivo tipos_de_datos.md
