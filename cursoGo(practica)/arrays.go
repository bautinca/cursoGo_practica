package main

import "fmt"

func main() {

	// Vamos a ver todo sobre los ARREGLOS o ARRAYS
	// nosotros habiamos dicho que un string es como un array de chars visto desde la perspectiva de C
	// Bueno es eso un array, un conjunto de elementos del mismo tipo, por ejemplo un array de enteros
	// un array de chars, un array de punteros, etc.

	// Vamos a hacer un array de 3 enteros
	var arr [3]int

	// Al imprimir este array sin defini obviamente esta todo inicializado en 0 a diferencia de C donde C
	// si las variables no estan definidas les asigna valores basura random
	fmt.Println(arr) // Efectivamente se imprime [0 0 0]

	// Asi como declaramos arrays tambien podemos declarar matrices
	var matriz [3][2]float64 // Seria una matriz de numeros flotantes de 3x2, 3 filas y 2 columans
	fmt.Println(matriz)

	// Ahora hasta ahora el arr y matriz esta todo inicializado en 0, podemos asignarle valores a ubicaciones especificas
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Tambien podemos acceder a sus indices, osea hasta ahora es todo igual que en otros lenguajes
	fmt.Println(arr[0])

	// Ahora para declarar y definir un array todo en la misma linea:
	vector := [3]int{5, 43, 10}
	fmt.Println(vector)

	// Tambien al igual que C no es necesario especificar el tamaño, aunque si es altamente recomendable que el espacio del array no sea una variable
	// es decir, que sea un arreglo ESTATICO si o si
	arreglo2 := [...]int{2, 4, 1, 49, 4, 2, 56, 632, 2, 353, 223}
	fmt.Println(arreglo2)

	// Logicamente con la funcion len() podemos sacar el tamaño de cualquier arreglo
	fmt.Println(len(arreglo2))

	// Ahora los arreglos tambien los podemos comparar!, por ejemplo:
	a := [3]int{1, 4, 2}
	b := [3]int{1, 8, 2}

	fmt.Println(a == b) // me devuelve false ya que es mentira que el array a es completamente igual al array de b
	fmt.Println(a != b) // Obviamente es true, los arrays son distintos
	fmt.Println(a == a) // Obviamente da true, el array a es exactamente igual al array a

}

// Ahora vamos al archivo slays.go
