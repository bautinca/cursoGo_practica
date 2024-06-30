package main

import "fmt"

func main() {
	// Vamos a ver como hacer un MAP, son los famosos diccionarios o hashes en
	// otros lenguajes, logicamente tenemos un <clave:valor>

	// Vamos a crear un map:
	x := make(map[string]string) // Como vemos el make() ademas para slices, sirve para crear maps, la clave:valor
	// seran ambas de tipo string, lo que va entre corchetes es la clave, y lo otro es el valor
	// Logicamente inicialmete el map va a estar vacio
	fmt.Println(x)

	// Tambien cuando declaramos un map le podemos definir su capacidad de entrada tambien con make:
	y := make(map[string]string, 2)
	fmt.Println(y)

	// Es decir, claramente los MAPS son DINAMICOS, y si es dinamico porque le tenemos que declarar una capacidad?
	// Pues obviamente para mejorar la performance del programa, es mucho mas conveniente darle una capacidad que sabemos que map no va a superar
	// Osea si ya de entrada le indico una capacidad que se que voy a usar al map es mas eficiente que inicialmente no declararle capacidad y por lo tanto GO va a tener
	// que ir haciendo sucecivos mallocs() para ir pidiendo mas memoria dinamica

	// Como le agrego el par <clave:valor> a un map? Como en cualquier otro lenguaje con los diccionarios:
	// obviamente la clave:valor tiene que ser ambos de tipo string en el caso de la variable 'x' e 'y'

	x["nombre"] = "Alejandro"
	x["edad"] = "29"
	fmt.Println(x)

	// Tambien para acceder a los valores de cada clave es tal cual identico que en python:
	fmt.Println(x["nombre"]) // Se imprimiria logicamente el valor de la calve "nombre", osea "Alejandro"

	// Tambien cuando declaramos un map() tambien lo podemos definir en la misma linea
	edades := map[string]int{
		"ana":       55,
		"rafael":    23,
		"manuel":    26,
		"alejandro": 29,
		"maria":     15,
	}

	fmt.Println(edades) // Logicamente se va a imprimir todo el map

	// Ahora, ¿Como eliminar elemento de un map? Usamos la funcion delete() donde le indicamos el map y la clave con su valor que queremos eliminar:
	delete(edades, "ana") // Se tendria que eliminar la clave "ana" con su respectivo valor 55
	fmt.Println(edades)

	// Por otro lado ¿Que pasaria si consultamos el valor de una clave que no existe? Bueno nos devolveria 0 (si el valor es int) o una clave nula (si el valor es string)
	// pero es interesante decir que NO DA ERROR
	fmt.Println("La edad de Pedro es:", edades["pedro"]) // Nos devuelve el entero 0

	// Logicamente al map se lo puede pensar como un diccionario en python, por lo tanto se le puede hacer las mismas operaciones matematicas, por ejempo:
	edades["pedro"]++     // Incrementa +1 el valor de la clave "pedro"
	edades["carlos"] += 2 // Incrementa +2 el valor de la clave "carlos"

	fmt.Println(edades) // Como vemos siempre que insertamos una clave nueva (como "pedro" y "carlos") se inicializan en 0,
	// por tanto es logico que al hacer ++ me quede 1 y al hacer += 2 me quede 2

	// Ahora algo importante es que entonces, si cada clave por mas que no exista va a tener un valor de 0
	// ¿Como me aseguro si una clave existe o no? Bueno como solucion a esto nosotros cuando consultamos por una clave nos va a dar un segundo
	// valor donde nos dice si ese valor existia o no, por ejemplo:

	// Claramente el nombre "pepe" en el diccionario no existe, por tanto:
	edad, ok := edades["pepe"]

	fmt.Printf("La edad de Pepe es %d, por existe? %t\n", edad, ok)

	// Entonces, como vemos todas las claves (existan o no en el map) nos va a dar un segundo valor oculto que es un booleano
	// donde justamente indica si la clave existe en el map o no

	// Ahora, al igual que los diccionarios en Python, los maps tienen un largo, y para sacar el largo es logicamente con len()
	fmt.Println("El largo del map 'edades' es:", len(edades))

	// Otra cosa importante de los maps es que LOS VALORES NO SON VARIABLES, por tanto, como no son variables no tienen una direccion
	// que les pertenezca, por tanto no pueden tener punteros que les apunten.
	// En GO los punteros se declaran igual que en C, con el &:

	//puntero := &edades["carlos"] // Esto nos daria error

	// Ahora bien ¿Como podemos recorrer un map? Bueno hay varias maneras pero una manera muy prolija es, como ya vimos, con el RANGE!
	for nombre, edad := range edades {
		fmt.Printf("El nombre es %s y su edad es %d\n", nombre, edad)
	}

}

// Ahora vamos al archivo funciones1.go
