package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Que es un slice? Es basicamente una parte continua de un array,
	// un slice vendria a ser como un arreglo dinamico, cuando querramos puede cambiar de tamaño pero al final justamente lo que engloba al slice es un array
	// Las slices se declaran asi:
	var j []int // Un slice no es necesario aclararle cuantos valores va a tener
	// Como se inicializa un slice vacio? Obviamente vacio, osea []
	fmt.Println(j)

	// Los slices son casi identicos a los arrays, en el sentido que tambien se pueden declarar y definir asi:
	x := []int{1, 2, 3}
	fmt.Println(x)

	// Para ver de que tipo de dato es 'x' podemos del paquete REFLECT usar la funcion TypeOf()
	fmt.Println(reflect.TypeOf(x)) // Si nos fijamos nos dice que es de tipo []int, osea un arreglo de enteros, osea un slice
	// porque al fin y al cabo un slice no deja de ser un arreglo ya que forma parte de un arreglo mas grande

	// Las operaciones que se les puede hacer a un slice es identico a un array, la unica diferencia es nunca a un slice se le dice cuantos elementos son ya que un slice
	// justamente es algo DINAMICO a diferencia de un array comun y corriente que es ESTATICO

	// Tambien para declarar un slice se puede usar la funcion make() donde le tenemos que pasar de que tipo queremos
	// que sea el slice y el largo inicial
	y := make([]int, 5)
	fmt.Println(y) // Obviamente se va a imprimir todo inicializado en 0 [0 0 0 0 0]
	// Ahora sacamos el tamaño del slice:
	fmt.Println(len(y))
	// Y la CAPACIDAD del slice ¿Cual es la diferencia entre un len() y una capacidad? Es que justamente el len indica
	// el tamaño exacto del slice en ese momento, pero la CAPACIDAD cap() indica la capacidad del array que en su interior tiene ese mismo slice!
	// Ya que justamente recordemos que dijimos que un slice es un arreglo pero que forma parte de otro arreglo mas grande
	fmt.Println(cap(y))

	// Si ejecutamos el programa nos va a decir que el len del slice 'y' es 5 pero su capacidad tambien! Ya que el arreglo que engloba
	// al slice es el propio slice!

	// Entonces como modificamos la capacidad de un slice? Bueno make() tiene un tercer parametro
	// que es para definir la CAPACIDAD de ese slice!
	k := make([]int, 5, 10)
	fmt.Println(cap(k)) // Logicamente la capacidad es 10

	// Ahora como el slice 'k' esta dentro de un array mas grande de capacidad 10 entonces el slice k se puede mover como
	// quiera dentro de ese array de capacidad 10

	// Ahora lo que vamos a hacer es declarar y definir un ARREGLO ESTATICO, para ello logicamente le tenemos que definir su longitud (o no)
	// pero si no le definimos la longitud de todos modos la longitud fija que le va a quedar es dependiendo de la cant. de elementos que le coloquemos cuando la definimos
	// por primera vez:
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // ARREGLO ESTATICO
	fmt.Println(ar)

	// Ahora vamos a crear 2 slices vacios
	var a, b []int
	fmt.Println(a)
	fmt.Println(b)

	// Ahora atencion aca, nosotros dijimos que los slices son partes de arrays no?
	// Entonces 'ar' sabemos que es un arreglo, y 'a' y 'b' slices puros
	// Por tanto lo que haremos es al slice 'a' asignarle un segmento de 'ar' y al slice 'b' otro segmento de 'ar'

	a = ar[2:5] // Lo que hacemos aca es que el slice 'a' sea desde el indice 2 hasta el indice 4 (inclusive) del array 'ar'
	// Entonces ahora logicamente la longitud (len) del slice 'a' ahora es 3!
	fmt.Println(a) // Se imprime logicamente [2 3 4] que es parte del array 'ar'

	// Ahora hacemos lo mismo pero con el slice 'b', queremos que sea parte tambien del array 'ar'
	b = ar[3:5]    // El slice 'b' es desde el indice 3 hasta el 4 (inclusive) del array 'ar'
	fmt.Println(b) // Logicamente se imprimiria una parte del array 'ar' que es [3 4] que seria el slice 'b'

	// Ahora si el slice 'a' y 'b' forman parte del array 'ar' eso quiere decir que si modifico
	// un elemento del slice 'a' tambien se modifica ese mismo elemento en el array 'ar'? Veamos...
	a[0] = 111 // Seria el indice 2 del array 'ar'
	// Verificamos...
	fmt.Println(a)
	fmt.Println(ar)
	// EFECTIVAMENTE!!! Como los slices siempre salen de un array eso quiere decir que estan directamente conectados
	// los slices del array de donde salen, es decir, los slices son como los hijos de un array estatico

	// Ahora, otra cosa a explicar es que tanto el slice 'a' como el slice 'b' comparten el elemento del indice 4 del array 'ar'
	// eso quiere decir que si modifico este elemento del indice 4 tambien se va a modificar en los slices que compartan ese elemento y el mismo array?
	// Veamos:

	b[1] = 444 // Seria indice 4 en array 'ar' e indice 2 en slice 'a'
	fmt.Println(a)
	fmt.Println(ar)
	fmt.Println(b)
	// EFECTIVAMENTE!!, entre slices de un mismo array estan directamente relacionados, es decir, si un slice sufre una modificacion
	// entonces el resto de los slices, y obviamente el array comparten dicha modificacion

	// Ahora cual seria la capacidad del slice 'a' y 'b'?
	fmt.Println(cap(a)) // 8
	fmt.Println(cap(b)) // 7

	// Nos da resultados distintos, porque? Si ambos no forman parte del mismo array 'ar'? Si
	// Pero la capacidad cuenta en adelante, es decir, no tiene en cuenta los elementos de atras del array 'ar'

	// Osea para entender mejor, supongamos que tenemos:

	// |_0_|_1_|_2_|_3_|_4_|_5_|_6_|_7_|_8_|_9_| --> Array 'ar'
	//		     ↑---a---↑
	//			     ↑-b-↑

	// Esto quiere decir que la capacidad del slice 'a' seria a desde el indice 2 de 'ar' en adelante, osea
	// la capacidad del slice 'a' es 8, y con la misma logica la capacidad del slice 'b' es del indice 3 del 'ar' en adelante
	// osea la capacidad de slice 'b' es 7

	// ¿Porque se tienen en cuenta solo los casilleros de adelante? Porque el slice crece hacia adelante! y no hacia atras! dentro de un array

	// ----------------------------------------------------------------
	fmt.Println("-----------------------------------------------------")

	// Ahora bien, como alargamos un slice? Es decir, como le agregamos elementos? con el famoso APPEND!
	// Primero lo que vamos a hacer es definir un slice:
	slice := make([]byte, 4, 10)
	fmt.Println(slice)
	// Su capacidad
	fmt.Println(cap(slice))

	// Ahora lo que vamos a hacer es definirlo:
	slice = []byte{'H', 'O', 'L', 'A'} // El slice es de tipo de dato byte = int8 y aca pongo chars? Si, entonces logicamente
	// lo que sucederia es que la 'H' en unicode 72, la 'O' el 79, y asi...
	fmt.Println(slice)

	// Como hacer para que imprima "HOLA"? bueno usando printf() que funciona tal cual igual que en C!
	fmt.Printf("%s\n", slice) // Aca el printf() funciona tal cual igual que en C, por ejemplo yo mismo le tengo que agregar el salto de linea

	// Ahora bien, cual es la diferencia entre sprintf() y printf() en GO? Bueno, el printf() funciona EXACTAMENTE IGUAL que en C, es decir, sirve directamente
	// para imprimir por default en la STDOUT, en cambio el sprintf() sirve para FORMATEAR UNA CADENA, es decir, darle forma a una cadena, asignarsela a una variable y luego
	// hacer con esta variable (cuyo valor esta formateado) lo que querramos, si queremos imprimimos, o lo que sea
	// en cambio el printf() es para directamente imprimir si o si en la STDOUT

	// Luego en cuanto a funcionamiento printf() y sprintf() funcionan EXACTAMENTE IGUAL, pero en finalidad son distintas

	// Es decir, printf() y sprintf() en GO funcionan exactamente igual que en C donde el sprintf() en C era para formatear una cadena y esta cadena pasarsela a un buffer ya formateada

	// Ahora lo que vamos a hacer es imprimir cada elemento del array 'slice', que recordar que todo slice es un array pero dinamico
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice x[%d]: %c\n", i, slice[i]) // Recordar que %c es para un char y %s es para un string
	}

	// Ahora, el 'slice' tiene una capacidad de 10, eso quiere decir que en teoria se le podria agregar mas elementos, osea hacer append
	// pero ojo, que tenga espacios libres para rellenar el slice no quiere decir que podramos hacer por ejemplo slice[5] = 'X', esto NO se puede hacer
	// ya que esto no es un append ya que si bien el slice se puede alargar aun el indice 5 del slice no esta habilitado

	// Entonces si yo quiero agregar un elemento al slice usamos el append()
	slice = append(slice, 'X')
	// Y listo!, verifiquemos:
	fmt.Printf("%s\n", slice) // Efectivamente se imprime HOLAX

	// Ahora cuanto seria la capacidad de slice?
	fmt.Println(cap(slice)) // Nos da 8! Como?

	// Bueno para entender bien esto de la CAPACIDAD veamos aca:
	fmt.Println("---------------------------------------")

	// Vamos a crear un slice completamente nuevo
	slice2 := make([]int, 0, 3)

	// Es un slice vacio, osea no tiene ningun elemento, por ende se deberia imprimir []
	fmt.Println(slice2)

	// Ahora para ver bien que es la capacidad lo que vamos a hacer es ir agregandole elementos
	// al slice (appends) e ir consultando por el len del slice y su capacidad:
	for i := 0; i < 15; i++ {
		slice2 = append(slice2, i)
		fmt.Printf("La longitud del slice es %d, y su capacidad es %d\n", len(slice2), cap(slice2))
	}

	/*
		Esto fue lo que nos imprimio este ultimo for:
		La longitud del slice es 1, y su capacidad es 3
		La longitud del slice es 2, y su capacidad es 3
		La longitud del slice es 3, y su capacidad es 3
		La longitud del slice es 4, y su capacidad es 6
		La longitud del slice es 5, y su capacidad es 6
		La longitud del slice es 6, y su capacidad es 6
		La longitud del slice es 7, y su capacidad es 12
		La longitud del slice es 8, y su capacidad es 12
		La longitud del slice es 9, y su capacidad es 12
		La longitud del slice es 10, y su capacidad es 12
		La longitud del slice es 11, y su capacidad es 12
		La longitud del slice es 12, y su capacidad es 12
		La longitud del slice es 13, y su capacidad es 24
		La longitud del slice es 14, y su capacidad es 24
		La longitud del slice es 15, y su capacidad es 24
	*/

	// Que quiere decir todo esto? Pues que en un principio el slice tenia una capacidad 3, por tanto le podiamos agregar efectivamente 3 elementos
	// Entonces logicamente cuando le agregamos los 3 primeros elementos nos va a decir que su capacidad es 3 ya que esa es su capacidad inicial

	// Ahora, lo interesante es que ¿Que pasa si agrego un 4to elemento? Logicamente estaria superando la capacidad de 3!, entonces lo que hace
	// GO es hacer el famoso MALLOC()!! DUPLICANDO LA CAPACIDAD ANTERIOR DEL SLICE PARA AGREGAR MAS ELEMENTOS!, entonces GO hace un nuevo MALLOC()
	// para duplicar la capacidad anterior asignandola a un nuevo bloque de memoria, copia los elementos del slice y los asigna al nuevo bloque de memoria con su nueva capacidad duplicada
	// y libera el bloque anterior!, es decir, miren todas las lineas de codigo que nos ahorramos en GO, esto hacerlo en C seria una locura

	// Entonces al alcanzar el 4to elemento sucede eso, GO hace un nuevo MALLOC() duplicando la capacidad anterior, se asigna un nuevo bloque de memoria, a este nuevo
	// bloque de memoria le ponemos los elementos que ya tenia el slice ya con su nueva capacidad duplicada y al bloque de memoria anterior con su capacidad vieja lo liberamos

	// Luego, obviamente al alcanzar el 7° elemento sucede lo mismo, como superamos la capacidad de 6 ahora hacemos un nuevo MALLOC() duplicando la capacidad anterior (6), por lo que
	// ahora tendriamos una capacidad de 12, copiamos los elementos que ya tenia el slice, lo colocamos en el nuevo bloque de memoria, y al bloque anterior de memoria con la capacidad 6 lo liberamos

	// Y asi sucesivamente... Siempre al alcanzar la capacidad maxima se realiza un nuevo malloc() duplicando la capacidad. y al anterior bloque de memoria obviamente GO lo libera

	fmt.Println("-------------------------------------")

	// Ahora vamos a ver como copiar el contenido de un SLICE en otro SLICE
	// Primero vamos a hacer 2 slices comunes y corrientes
	origen := []int{1, 2, 3}
	fmt.Printf("El largo del slice es %d, y su capacidad es %d\n", len(origen), cap(origen))

	destino := []int{3, 4, 5}
	fmt.Printf("El largo del slice es %d, y su capacidad es %d\n", len(destino), cap(destino))

	// Logicamente la capacidad de ambos es equivalente al largo inicial, luego cuando les haga un append ahi es cuando sucedera el malloc() duplicando la capacidad de ambos a 6

	// Bueno, ahora voy a hacer un copy() donde le tengo que pasar 2 slices, yo quiero que se copie el slice 'origen'
	// y se pegue en el slice 'destino', por tanto:
	copy(destino, origen)
	fmt.Println(origen, destino) // copy(lo_q_se_copia, donde_se_pega)

	// Ahora, esta fue una copia dentro de todo 'comoda' ya que ambos slices tenian una longitud de 3 y una capacidad de 3
	// Que pasa si ahora son distintos en ese sentido? Veamos:

	origen2 := []int{1, 2, 3}
	destino2 := make([]int, 1, 2)

	// La logica nos dice que si copiamos el contenido de origen2 y lo pegamos en destino2 entonces
	// logicamente superariamos la capacidad inicial 2, por lo tanto su capacidad pasaria a 4 con un len de 3 obviamente
	// Pero en realidad NO SUCEDE ESTO, lo que sucede es que el copy() toma los espacios disponibles que tiene el destino sin importar que su capacidad sea de 100 por ejemplo
	copy(destino2, origen2)
	fmt.Printf("La longitud de destino2 es %d, y su capacidad es de %d\n", len(destino2), cap(destino2))
	fmt.Println(destino2)
	// Entonces, lo que sucedio es que como destino2 tiene una longitud de 1 entonces solamente se copia el primer elemento
	// de origen2 a destino2, y su capacidad obviamente sigue siendo 2, es decir, el copy() no hace que se supere la capacidad, sino que tiene en cuenta el len del slice para hacer la copia y el pasteo
	// Entonces en este caso solo se copiaria el 1 del origen2

	// En cambio si el slice origen es menor al slice destino entonces la copia del origen es completa, y se pega completamente en el destino
	// pero el slice destino ademas va a conservar sus valores originales en los espacios que quedan, es decir:
	origen3 := []int{1, 2}
	destino3 := []int{3, 4, 5}

	copy(destino3, origen3)
	// Lo que deberia suceder es que destino3 quede {1,2,5} y su capacidad obviamente permanece en 3 ya que el copy() nunca modifica la capacidad
	fmt.Println(destino3)
}

// Ahora vamos al archivo range.go
