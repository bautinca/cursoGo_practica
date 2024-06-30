package main

import (
	"fmt"
)

// Creamos un struct
type storyPage struct {
	text     string
	nextPage *storyPage // Aca tenemos como atributo otro struct storyPage!, es decir, son como distintas paginas q estan enlazadas y forman un LIBRO!
}

// Ahora creamos una funcion para reproducir el libro pagina por pagina, pero
// esta funcion sera un metodo del struct storyPage
func (page *storyPage) playStory() {

	// Al meternos con la recursividad tenemos que atajar el caso base que es cuando
	// la actual pagina es nil
	if page == nil {
		return
	}

	// Reproducimos el texto de esa pagina
	fmt.Println(page.text)

	// Luego recursivamente llamamos a la pag. siguiente
	page.nextPage.playStory()
}

// Vamos a crear otro metodo asociado a struct storyPage que lo que hace es agregar una pagina al libro
func (page *storyPage) addToEnd(text string) {
	// El texto lo vamos a transformar en una pagina
	pageToAdd := storyPage{text, nil}
	// Para ello tendremos que iterar hasta encontrar la pagina cuyo atributo nextPage sea nil, y ahi agregamos
	// la pagina nueva
	for page.nextPage != nil {
		page = page.nextPage
	}
	// Si salimos del for es porque efectivamente llegamos a la pagina cuyo next es nil, por lo tanto
	// ahora si agregamos la pagina nueva
	page.nextPage = &pageToAdd
}

// Ahora vamos a crear otro metodo donde lo que hace es tambien agregar una pagina pero en el medio de
// un libro, para ello le tenemos que pasar la pagina donde luego de esa pagina queremos agregar nuestra pagina
func (page *storyPage) addAfter(text string) {
	newPage := storyPage{text, page.nextPage}
	// Ahora tenemos que cambiar el nextPage de la pagina actual que apunta hacia la pagina que creamos
	page.nextPage = &newPage
}

func main() {

	// Sabiendo el anterior struct lo que vamos a definir son una serie de paginas enlazadas formando un libro
	page1 := storyPage{"It was a dark and stormy night.", nil}

	// Ahora la pagina 2 tendra logicamente otro texto
	page1.addToEnd("You are alone, and you need to find the sacred helmet before the bad guys do")

	// Ahora la pagina 3
	page1.addToEnd("You see a troll ahead")

	// Vamos a agregar una pagina despues de la pagina 1
	page1.addAfter("Testing AddAfter")

	// Reproducimos las paginas del libro
	page1.playStory()

	// Ahora bien, vamos a ver las diferencias entre los siguientes conceptos:

	//	- FUNCIONES: Tiene un valor de retorno y ejecuta comandos
	//	- PROCEDIMIENTOS: Ejecuta tambien comandos pero no tiene valor de retorno
	//	- METODOS: Son funciones pero ligadas a structs
}
