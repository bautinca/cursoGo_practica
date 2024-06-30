package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Vamos a ver GRAFOS

// Nosotros podemos tener un nodo y a los costados del nodo varias direcciones, por ejemplo podemos
// estar en una habitacion y podemos ir a 4 habitaciones contiguas distintas, o 5, o 6, etc.

// Pero a su vez cada habitacion contigua puede tener mas habitaciones para poder visitar

// A su vez si estamos en una habitacion podemos ir a otra habitacion y de esa otra habitacion volver a la misma habitacion

// Tambien podemos hacer un recorrido ciclico por habitaciones, es decir, 1->2->3->1->2->3, y asi, visitar ciclicamente
// ciertas habitaciones

// Ahora otro struct relacionado a las elecciones de camino del usuario, cada choice tendra el puntero a otro choice
// osea claramente cada eleccion lleva a otra eleccion que es la ruta que siguio el usuario

// Tambien vamos a tener otro atributo descripcion de las elecciones que tomo el usuario

// Luego el comando es para que el usuario ingrese que camino tomar, o que arista tomar

// Por ultimo cada eleccion nos va a llevar a un nuevo struct storyNode, osea el atributo 'nextNode'
type choices struct {
	cmd         string     // El cmd, osea comando es el comando que el usuario ingresa para tomar una eleccion
	description string     // Representa una breve descripcion de cada eleccion (choice) que tomo el usuario
	nextNode    *storyNode // Cada eleccion del usuario nos lleva a un siguiente nodo
	nextChoice  *choices   // Cada eleccion es un puntero a la sig. eleccion
}

// Ahora bien, claramente en los grafos tenemos 2 elementos; NODOS Y ARISTAS. Entonces lo que vamos a hacer
// es hacer la representacion de un nodo, donde cada nodo tendra un texto y ademas otro atributo que es 'choices' que son
// basicamente las elecciones de recorrido del usuario hasta llegar al actual nodo, obviamente 'choices' sera otro struct
type storyNode struct {
	text    string   // Texto de cada nodo
	choices *choices // Las elecciones del usuario hasta llegar al nodo
}

// Ahora creamos un metodo que lo que hara sera agregar elecciones a un nodo, asi un nodo puede tener como 1 eleccion, 2, 3, 100, etc.
func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	// Con los parametros que ingreso el usuario creamos una nueva eleccion para el nodo
	choice := choices{cmd, description, nextNode, nil}

	// Si el nodo actual no tiene eleccion entonces se la agregamos su primer eleccion
	if node.choices == nil {
		node.choices = &choice

		// Si el nodo ya tiene elecciones entonces le agregamos esta nueva eleccion al final de todo de las elecciones
		// que tiene el nodo, osea porque cada nodo tiene una lista enlazada de elecciones posibles
	} else {
		// La eleccion actual la asignamos en una nueva variable
		currentChoice := node.choices
		// Agregamos la eleccion al final de todo de la lista enlazada de choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		// Si salimos del for es porque efectivamente llegamos a la eleccion cuya siguiente eleccion
		// esta vacia, por tanto ahi agregamos la eleccion que ingreso el usuario como parametros
		currentChoice.nextChoice = &choice
	}
}

// Ahora creamos otro metodo
func (node *storyNode) render() {
	// Simplemente lo que hace es imprimir el texto del nodo
	fmt.Println(node.text)
	// Ahora asignaremos a una variable la eleccion actual
	currentChoice := node.choices
	// Ahora lo que haremos sera iterar por todas las elecciones que tiene el nodo
	for currentChoice != nil {
		// Vamos a imprimir el comando de cada eleccion y su descripcion
		fmt.Println(currentChoice.cmd, ":", currentChoice.description)
		// Pasamos a la siguiente eleccion del nodo para ir iterando por todas las posibles elecciones que tiene el nodo
		currentChoice = currentChoice.nextChoice
	}
}

// Ahora creamos otro metodo que lo que haga sera ejecutar el comando
func (node *storyNode) executeCmd(cmd string) *storyNode {
	// A la eleccion actual del nodo la asignamos a una variable
	currentChoice := node.choices
	// Iteramos por todas las elecciones del nodo que tiene
	for currentChoice != nil {
		// Si el comando que ingreso el usuario coincide con el comando de una eleccion del nodo
		// entonces entramos al siguiente if, pero ademas lo que haremos sera aplicar ToLower() que es para pasar todo a minusculas una cadena
		// asi justamente ingoramos las mayusculas
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			// Si el comando que ingreso el usuario coincide con el comando de la eleccion entonces
			// retornamos el siguiente nodo al que lleva ese comando
			return currentChoice.nextNode
		}
		// Si no entramos al if es porque el comando que ingreso el usuario no coincide con el comando de ninguna eleccion
		// por tanto pasamos a la sig. eleccion
		currentChoice = currentChoice.nextChoice
	}
	// Si salimos del for es porque el comando que ingreso el usuario es un comando cualquiera que no coincide
	// con el comando de eleccion de ninguno, por tanto le comunicamos por STDOUT al usuario que ingreso cualquier cosa
	fmt.Println("Sorry, i didn't understand that.")
	// Y retornamos el mismo nodo
	return node
}

// Vamos a declarar una variable global que es un puntero a un scanner donde la usaremos luego
var scanner *bufio.Scanner

// Por ultimo hacemos el ultimo metodo que lo que hace es reproducir toda la historia
func (node *storyNode) play() {
	// Primero imprimimos el texto del actual nodo
	node.render()
	// Ahora si el nodo actual tiene elecciones entonces entramos al if
	if node.choices != nil {
		// Lo que hacemos es scannear el STDIN esperando a q el usuario ingrse algo
		scanner.Scan()
		// Ahora lo que hacemos es que lo que ingrese el usuario, que sera un COMANDO, bueno con este comando ejecutaremos
		// el nodo actual que es para tomar una eleccion y aca usaremos el metodo que implementamos executeCmd()

		// Ademas sera algo recursivo, ya que esto me devuevle el nodo siguiente dependiendo de la eleccion que tomo el usuario ingresando
		// el comando, por tanto al nodo siguiente que me devuelve le aplicamos nuevamente play() y asi siempre le pedira comandos al usuario
		node.executeCmd(scanner.Text()).play()
	}

}

// Por fin el main...
func main() {
	// Establecemos el STDIN como scanner
	scanner = bufio.NewScanner(os.Stdin)

	// Creamos un primer nodo
	start := storyNode{text: `
	You are in large chamber, deep underground.
	You see three passages leading out. A north passage leads into darkness.
	To the south, a passage appears to head upward. The eastern passages appears
	flat and well traveled`}

	// Ahora el primer nodo que creamos claramente no tiene elecciones, por lo tanto
	// le agregaremos elecciones o habitaciones por asi decirlo
	// Se puede tomar el camino de una habitacion oscura que sera otro nodo
	darkRoom := storyNode{text: "It is pitch black. Tou cannot see a thing."}

	// Otra habitacion
	darkRoomLit := storyNode{text: "The dark passae is now a lit by your lantern. You can continue north or head back south"}

	// Otra mas
	grue := storyNode{text: "While stumbling around in the darkness, you are eaten by a grue"}

	// Otra mas
	trap := storyNode{text: "You head down the well traveled path when suddenly a trap door opens and you fall into a pit"}

	// Otra mas
	treasure := storyNode{text: "Your arrive at a small chamber, filled with a trasure!"}

	// Ahora empezamos agregando elecciones ya tenemos todos los nodos
	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	// Ya le agregamos 3 habitaciones contiguas a start, ahora a otra habitacion, por ejemplo darkRoom
	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on lantern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	// Arrancamos el recorrido
	start.play()

	fmt.Println()
	fmt.Println("The End!")

}
