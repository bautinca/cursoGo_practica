package main

import (
	"bufio"
	"fmt"
	"os"
)

// Creamos un struct Nodo donde dependiendo de lo que ingrese el usuario se agregara tal pagina o tal otra el libro
// Es decir, es un libro que se formara dependiendo de las decisiones que tome el usuario
type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

// Creamos un metodo que lo que hace es imprimir la pagina y coloca como scanner
// el STDIN para luego tomar las decisiones que toma el usuario y de acuerdo a esas decisiones imprimir las paginas siguientes
func (node *storyNode) play() {
	// Antes que nada imprimimos la pagina actual
	fmt.Println(node.text)

	// Ahora preguntamos si no hay ninguna pagina por si el usuario ingresa 'yes' o 'no' entonces finalizamos
	if node.noPath == nil && node.yesPath == nil {
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// Ahora entramos en un for true donde esperamos a que el usuario siempre ingrese 'yes' o 'no' y de acuerdo
	// a su decision que pagina siguiente imprimir
	for {
		// Esperamos que el usuario ingrese algo en STDIN
		scanner.Scan()
		// Lo guardamos en una variable
		answer := scanner.Text()

		// Ahora vamos a tener 3 casos, si el usuario ingreso 'yes', si ingreso 'no' o si ingreso algo distinto a esos 2 casos anteriores
		if answer == "yes" {
			// Si el usuario puso "yes" entonces logicamente usaremos recursividad, donde iriamos
			// a la sig. pagina que esta en el atributo yesPath
			node.yesPath.play()
			break

		} else if answer == "no" {
			// Aca tambien usariamos recursividad pero para la pagina ligada al noPath
			node.noPath.play()
			break

		} else {
			// Aca si el usuario ingreso cualquier otra cosa que no sea 'yes' o 'no' repetimos el ciclo
			fmt.Println("Esa respuesta no es una opcion, por favor ingrese 'yes' o 'no'")
		}
	}
}

// Ahora creamos otro metodo que lo que hace es recorrer por todos los nodos e imprimir el texto
// Primero hacemos la accion, luego llamamos recursivamente un nodo izquierdo (tomamos como yespath el nodo izquierdo)
// y luego recursivamente el nodo derecho (nopath es el nodo derecho). Esto lo vimos en Algoritmos 2 y es un recorrido
// tipo PREORDER DEL ARBOL BINARIO!!, en un recorrido de este estilo primero se hace la accion del nodo actual, luego visitar
// al nodo izq recursivamente y luego recursivamente visitar al nodo derecho, es PREORDER porque es 'pre', osea la accion antes de visitar los nodos hijos
func (node *storyNode) printStory() {
	fmt.Println(node.text)
	if node.yesPath != nil {
		node.yesPath.printStory()
	}
	if node.noPath != nil {
		node.noPath.printStory()
	}

}

func main() {

	// Basicamente sera un arbol de decisiones donde tenemos la primer pagina suprema que es el nodo padre
	// ariba de todos, luego dependiendo de que si el usuario ingresa 'yes' o 'no' la pagina se agregara a izquierda o derecha del nodo

	// Entonces este es el primer nodo del arbol arriba de todo, el nodo raiz
	root := storyNode{"You are at the entrance to a dark cave. Do you want to go in the cave?", nil, nil}

	// Ahora ccreamos 2 historias paralelas, uno que el usuario gano y otro que el usuario perdio
	winning := storyNode{"You have won!", nil, nil}
	losing := storyNode{"You have lost!", nil, nil}

	// Ahora agregamos estos 2 ultimos nodos al nodo raiz, uno como yes y otro como no
	root.yesPath = &losing
	root.noPath = &winning

	// Ahora le damos a play para reproducir la historia donde tenemos un arbol de decisiones por el usuario, si el usuario ingresa
	// a la cueva entonces pierde y si decide no ingresar entonces gana
	root.play()

	// Ahora imprimimos todos los nodos sus textos
	root.printStory()

}
