package main

import "fmt"

// Con interfaces nos referimos a como diagramar o estructurar las cosas, ya tenemos los structs con sus respectivos
// metodos ahora tenemos que hacerlo todo ordenado

// Primero vamos a crear un struct Persona
type Persona struct {
	nombre string
	email  string
	edad   int
}

// Ahora un metodo de Persona para  extraer su nombre
func (p Persona) Nombre() string {
	return p.nombre
}

// Y otro metodo para extraer su email
func (p Persona) Email() string {
	return p.email
}

// Ahora aca declaramos una funcion comun y corriente que lo que hace es recibir una Persona y hacer
// que se 'presente', osea:
func Presentarse(p Persona) {
	fmt.Println("Nombre:", p.Nombre())
	fmt.Println("Email:", p.Email())
}

// Ahora como podemos hacer de aplicar la funcion Presentarse() para otros structs y no solo para Persona?
// Aca lo que haremos ahora es un struct mas
type Moderador struct {
	Persona
	Foro string
}

// Metodo de Moderador que lo que hace es comunicar que se abrira el foro
func (m Moderador) AbrirForo() {
	fmt.Println("Abrir el foro")
}

// Otro metodo que comunica que se cerrara el foro
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar el foro")
}

// Y ahora otro struct Administrador
type Administrador struct {
	Persona
	seleccion string
}

// Un metodo de Administrador que lo que hace es abrir un foro
func (a Administrador) CrearForo() {
	fmt.Println("Abrir el foro")
}

// Y otro metodo de Administrador que lo que hace es cerrar un foro
func (a Administrador) CerrarForo() {
	fmt.Println("Cerrar el foro")
}

// Ahora que pasa si digo que la funcion comun y corriente Presentarse() quiero que ademas del struct Persona
// quiero que maneje al Moderador y Administrador?

// Entonces deberia crear 2 funciones mas una para que se presenten los Administradores y otro para los Moderadores

// Presentarse para los Moderadores...
func PresentarseM(m Moderador) {
	fmt.Println("Nombre:", m.Nombre())
	fmt.Println("Email:", m.Email())
}

// Recordar que los metodos Nombre() y Email() eran para el struct Persona, pero como dentro del struct Moderador tiene un atributo que es Persona
// entonces tambien acepta los metodos Nombre y Email

// Y un presentarse para Administradores
func PresentarseA(a Administrador) {
	fmt.Println("Nombre:", a.Nombre())
	fmt.Println("Email:", a.Email())
}

// Ahora vamos a presentar la instruccion INTERFACE ya que si nos fijamos PresentarseA(), PresentarseM() y Presentarse() hacen exactamente
// lo mismo, la unica diferencia es que reciben distinto tipo de argumento, uno recibe un Persona, otro un Administrador y otro un Moderador
// Entonces lo malo es que tenemos codigo repetido, entonces no es posible crear un Presentarse() que pueda recibir una Persona, Administrador o Moderador?
// Bueno la solucion esta en INTERFACE

type Usuario interface {
	Nombre() string // Se pone el metodo Nombre() que va a implementar y devuelve un string
	Email() string  // Y aca el otro metodo que devuelve un string diremos
}

// OJO, NO ES UN STRUCT,
// Entonces lo que voy a hacer ahora a la funcion Presentarse() es que originalmente era solo de Persona, ahora lo que voy a hacer
// es que la funcion Presentarse() reciba la interfaz Usuario, para ello crearemos la funcionPresentarseTodos()

func PresentarseTodos(p Usuario) {
	fmt.Println("Nombre:", p.Nombre())
	fmt.Println("Email:", p.Email())
}

// Y la funcion main
func main() {
	// Creamos una Persona
	alejandro := Persona{"Alejandro", "a@gmail.com", 29}
	Presentarse(alejandro)

	// Creamos un Moderador
	juan := Moderador{Persona{"Juan", "j@gmail.com", 46}, "Juegos"}

	// Creamos un Administrador
	pedro := Administrador{Persona{"Pedro", "p@gmail.com", 25}, "PC"}

	// Ahora presentamos a juan y a pedro
	PresentarseM(juan)
	PresentarseA(pedro)

	// Ahora lo cierto es que hacer PresetarseM() y PresentarseA() esta verdaderamente mal porque rompemos uno de los fundamentos
	// de la POO que es no repetirnos, es decir, lo ideal seria que la funcion Presentarse() sea aplicable para Moderadores y Administradores

	// Si nos fijamos Presentarse(), PresentarseM() y PresentarseA() hacen EXACTAMENTE LO MISMO, la unica diferencia es el tipo de parametro que reciben
	// uno recibe una Persona, otro recibe un Moderador y otro un Administrador

	// No podemos hacer que Presentarse() pueda recibir Persona o Moderador o Administrador? Para eso existe la instruccion INTERFACE{}
	// Ahora vamos a usar la funcion PresentarseTodos() que recibe la INTERFAZ USUARIO

	// Vamos a llamar a PresentarseTodos() para un alejandro(Persona), juan(Moderador) y pedro(Administrador)
	fmt.Println("-------------------------------------------")
	PresentarseTodos(alejandro)
	PresentarseTodos(juan)
	PresentarseTodos(pedro)

	// Y listo! Funciono bien!

	// Ahora bien, las interfaces, al igual que los structs, son tipos de datos, por lo tanto podemos declarar
	// variables del tipo interfaz Usuario!:
	var i Usuario
	// Que valor puede tener 'i'? Logicamente un struct donde sea compatible la interfaz Usuario, por ejemplo
	// 'i' puede equivaler a un 'alejandro' ya que 'alejandro' tiene los metodos Nombre y Email() tal cual como tiene la interfaz Usuario
	i = alejandro // (tipo Persona)
	// Ahora vamos a imprimir
	fmt.Println("----------------------------------------------")
	fmt.Println(i)
	// Tambien le podemos sacar su Nombre y Email
	fmt.Println("El nombre es:", i.Nombre())
	fmt.Println("El mail es:", i.Email())

	// Tambien logicamente 'i' puede equivaler a un tipo de dato Moderador
	i = juan
	fmt.Println(i)
	fmt.Println("El nombre es:", i.Nombre())
	fmt.Println("El mail es:", i.Email())

	// Y logicamente tambien 'i' puede equivaler a un tipo de dato Administrador ya que tambien tiene los metodos Nombre() y Email()
	i = pedro
	fmt.Println(i)
	fmt.Println("El nombre es:", i.Nombre())
	fmt.Println("El mail es:", i.Email())

}

// Ahora vamos al archivo errores.go
