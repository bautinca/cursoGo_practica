package main

import (
	"fmt"
	"math"
)

// Ahora vamos a hablar de los METODOS que son basicamente funciones pero exclusivas de las ESTRUCTURAS que definamos
// Osea los metodos existen para ir alterando el estado de las estructuras

// Primero lo que hacemos es definir una estructura
type Rectangulo struct {
	ancho, alto float64
}

// Voy a crear otro struct
type Circulo struct {
	radio float64
}

// Ahora definimos un metodo llamado area() que es para sacar el area de la estructura Rectangulo
// El metodo siempre recibe el TDA, luego colocamos el nombre del metodo, osea area() y por ultimo lo que devuelve que seria un float64
func (r Rectangulo) area() float64 {
	return r.ancho * r.alto
}

// Voy a crear otro metodo que se llama igual que el anterior, pero es para calcular areas de Circulos
func (c Circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}

// Ahora vamos a crear otro metodo pero de Rectangulo, lo que hace es que se llama inc() y recibe 1 parametro float64 y devuelve un Rectangulo modificado osea...
func (r Rectangulo) inc(i float64) Rectangulo {
	// Lo que hace el metodo es multiplicar el ancho y alto de 'r' por ese float64 'i' y devuelve ese mismo rectangulo modificado
	nuevoRectangulo := Rectangulo{r.ancho * i, r.alto * i}
	return nuevoRectangulo
}

// Este metodo es inc2() que es lo mismo que inc() pero optimizado ya que modifica INPLACE al struct Rectangulo y no devuelve uno nuevo como inc()
// Para que modifique INPLACE a un rectangulo entonces tiene que recibir EL PUNTERO A ESE RECTANGULO QUE QUIERO MODIFICAR!
func (r *Rectangulo) inc2(i float64) { // Logicamente no retorna nada ya que la modificacion la hace inplace
	r.ancho *= i
	r.alto *= i
}

func main() {

	// Ahora vamos a declarar y definir 2 Rectangulos
	r1 := Rectangulo{12, 2}
	r2 := Rectangulo{9, 4}

	// Ahora vamos a sacar sus areas:
	fmt.Println("El area de r1 es:", r1.area())
	fmt.Println("El area de r2 es:", r2.area())

	// Declaramos y definimos 2 Circulos:
	c1 := Circulo{10}
	c2 := Circulo{25}

	// Calculamos sus areas
	fmt.Println("El area de c1 es:", c1.area())
	fmt.Println("El area de c2 es:", c2.area())

	// Ahora lo que vamos a hacer es usar el metodo inc() para Rectangulo
	fmt.Println("El rectangulo 1 es:", r1)
	// Incrementamos el r1
	r1 = r1.inc(10)
	fmt.Println("El rectangulo 1 incrementado quedo:", r1)

	// Ahora bien, aca nos llama la atencion algo, y es que NO ESTAMOS MODIFICANDO EL MISMO R1, sino que el metodo inc() CREA
	// UN NUEVO RECTANGULO y este nuevo rectangulo reemplaza a r1, pero esto estaria mal en teoria, ya que no estamos modificando INPLACE a r1
	// Entonces como podemos modificar inplace un struct? USANDO PUNTEROS!, y lo veremos en el metodo inc2()

	// Ahora vamos a testear fun2()
	fmt.Println("El rectangulo 2 es:", r2)
	r2.inc2(25)
	fmt.Println("El rectangulo 2 incrementado quedo:", r2)
}

// Ahora vamos a estructuras4.go
