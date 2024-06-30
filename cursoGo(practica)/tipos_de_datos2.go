/*
Ahora vamos a hablar TODO SOBRE LOS TIPOS DE DATOS ENTEROS como vimos antes
en el archivo Markdown

Vamos a ver uno por uno todos los tipos de datos enteros
*/

package main

// Tambien importaremos el paquete unsafe para usar el sizeof
import (
	"fmt"
	"unsafe"
)

func main() {
	var entero0 int8 // Este es el tipo de dato entero int8, logicamente
	// es un tipo de dato que es un entero con signo y ocupa 8 bits
	// Por tanto, como contamos con 8 bits para representar cualquier numero con signo
	// entonces (2^8) = 256 numeros distintos podemos representar
	// pero como representamos con signo entonces lo dividimos por la mitad, una mitad para representar
	// los numeros negativos y la otra mitad para los numeros positivos.
	// Entonces el int8 representa numeros del intervalo [-128, 127] (logicamente hay que tener ademas
	// en cuenta el 0)

	// Ahora vamos a calcular cuanto pesa con la funcion sizeof() que tambien la podemos usar en GO ya que importamos
	// el paquete unsafe, logicamente nos va a decir que lo que pesa entero0 es 8 bits = 1 byte, y recordar que sizeof() nos da el resultado
	// en bytes, osea en este caso nos deberia imprimir 1

	// Notar como hacemos el Println() para imprimir un texto y el valor de una variable, es mucho mas flexible que C en ese sentido
	fmt.Println("El tamaño del int8 es:", unsafe.Sizeof(entero0))

	// Ahora vamos a ver otro entero, este entero ocupa 16 bits = 2 bytes, por lo tanto es logico que representa
	// un mayor rango de valores enteros ya que tenemos 2 bytes disponibles, por tanto es logico que pesa 2 bytesd
	var entero1 int16
	// Imprimimos cuanto pesa
	fmt.Println("El tamaño del int16 es:", unsafe.Sizeof(entero1)) // Esta es otra manera de tambien usar Println()

	// Tambien tenemos el entero de 32 bits! obviamente representa un rango muchisimo mas amplio de numeros
	var entero2 int32
	fmt.Println("El tamaño del int32 es:", unsafe.Sizeof(entero2))

	// Por ultimo el mayor de todos es el que ocupa 64 bits! osea una barbaridad, en total son 8 bytes! lo que va a ocupar esta variable
	var entero3 int64
	fmt.Println("El tamaño del int64 es:", unsafe.Sizeof(entero3))

	// Tambien tenemos los tipos de datos que representan enteros pero SIN SIGNO y son:

	// El primero es el tipo de dato uint8 (donde la u significa unsigned, es decir, no tiene signo) y ocupa logicamente 8 bits
	// Ademas como (2^8) = 256 eso quiere decir que representamos 256 numeros distintos pero POSITIVOS, osea [0, 255] son los numeros que representa
	var entero4 uint8
	fmt.Println("El tamaño del uint8 es:", unsafe.Sizeof(entero4))

	// Y logicamente para no hacer todo repetitivo tambien tenemos el entero sin signo de 16, 32 y 64 bits:
	var entero5 uint16
	var entero6 uint32
	var entero7 uint64

	fmt.Println("El tamaño del uint16 es:", unsafe.Sizeof(entero5))
	fmt.Println("El tamaño del uint32 es:", unsafe.Sizeof(entero6))
	fmt.Println("El tamaño del uint64 es:", unsafe.Sizeof(entero7))

	// Tambien hay otros tipos de datos que no tiene C, o por lo menos en su entadar, que son byte y rune
	// donde byte sabemos que claramente son 8 bits, por lo tanto el tipo de dato byte = uint8
	// Por otro lado el tipo de dato rune = int32

	var entero8 byte
	var entero9 rune

	fmt.Println("El tamaño del byte es:", unsafe.Sizeof(entero8))
	fmt.Println("El tamaño del rune es:", unsafe.Sizeof(entero9))

	// Luego tenemos los que dependen de nuestra arquitectura del computador, logicamente en esta categoria esta INT!, donde lo que pesa depende
	// de la arq. del computador, pero casi siempre pesa 8 bytes en sistemas de 64 bits

	var entero10 int  // En sistemas de 64 bits pesa casi siempre 8 bytes
	var entero11 uint // Es como el int pero sin signo, en sistemas de 64 bits pesa casi siempre 8 bytes

	fmt.Println("El tamaño del int es:", unsafe.Sizeof(entero10))
	fmt.Println("El tamaño del uint es:", unsafe.Sizeof(entero11))

	// Ahora, como se hacen los CASTEOS? Es decir, que una variable pase de un tipo de dato a otro tipo de dato
	// Lo que vamos a hacer es declarar y definir 2 variables de 2 tipos distintos
	var valor1 int32 = 25
	var valor2 int64 = 85

	// Ahora sabemos que claramente son 2 tipos de datos distintos ¿Los podemos sumar?
	// fmt.Println(valor1 + valor2) // Esto NO SE PUEDE HACER, ya que al igual que C, son 2 tipos de datos distintos

	// Es por ello que ahora para que se sumen los vamos a CASTEAR, vamos a hacer que el tipo de dato de valor2 sea int32 en lugar de int64.
	// Y decidimos pasar este porque es mas optimo, porque ahora valor2 en lugar de ocupar 8 bytes va a pasar a ocupar 4 bytes
	fmt.Println(valor1 + int32(valor2))

	// Ahora si hacemos la suma entre un INT32 y un RUNE esto logicamente va a funcionar ya que dijimos que RUNE = INT32
	var valor3 int32 = 10
	var valor4 rune = 15

	fmt.Println(valor3 + valor4) // Esto si se puede hacer

	// ¿Y si sumo INT con INT32? logicamente no voy a poder, ya que en mi arq el INT pesa 64 bits por tanto lo voy a tener
	// que castear
	var valor5 int = 3
	var valor6 int32 = 5

	fmt.Println(int32(valor5) + valor6)
	// O tambien podriamos hacer
	fmt.Println(valor5 + int(valor6))
}

// Ahora vamos a ir al archivo operadores.go
