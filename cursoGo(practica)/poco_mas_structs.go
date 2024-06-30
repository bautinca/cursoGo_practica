package main

import "fmt"

// Vamos a ver un poco mas de los structs en GO
// Vamos a crear un struct Position
type Position struct {
	// Atributos
	x float32
	y float32
}

// Hacemos otro struct donde uno de sus atributos es otro struct
type badGuy struct {
	name   string
	health int
	pos    Position
}

// Creamos una funcion que incrementa +1 un numero que recibe, para ello recibe un puntero a un int logicamente
func addOne(num *float32) {
	*num++
}

// Tambien vamos a crear una funcion donde nos diga donde se encuentra badGuy
func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Printf("(%f, %f)\n", x, y)
}

// Ahora el main...
func main() {
	// Declaramos una variable 'p' que es Position
	var p Position

	// Asignamos valor al atributo 'x'
	p.x = 5
	// Asignamos valor al atributo 'y'
	p.y = 10

	// Imprimimos el atributo 'x' e 'y' de 'p' que es un struct Position
	fmt.Println(p.x)
	fmt.Println(p.y)

	// Tambien podemos declarar y definir un Position todo en la misma linea recordemos
	o := Position{4, 2}

	// Imprimimos sus atributos
	fmt.Println(o.x)
	fmt.Println(o.y)

	//-------------------------------------

	// Declaramos y definimos otro struct
	badguy := badGuy{"Jabba The Hut", 100, p}

	// Imprimimos asi nomas el struct badGuy 'badguy'
	fmt.Println(badguy)

	// Ahora consultamos a donde esta badGuy
	whereIsBadGuy(badguy)

	// Vamos a incrementar +1 la coordenada x de badguy
	addOne(&badguy.pos.x)

	// Ahora imprimimos las coordenadas nuevamente
	whereIsBadGuy(badguy)

	return
}
