/*
Bueno, como vemos en GO se colocan comentarios IGUAL que en C, ademas
tambien para colocar comentarios de una sola linea podemos colocar //

Primero que todo vamos a ver como IMPRIMIR cosas en la STDOUT, claramente sabemos que en C es printf()
y en Python es print(), ¿Como sera en GO?

Todo programa en GO debe tener un paquete llamado MAIN, que es el que se ejecuta en primer instancia
Siempre que nosotros arranquemos a hacer un programa en GO siempre al principio lleva un PACKAGE, en este caso lo que diremos es que todo el archivo primer_programa.go pertenece
al paquete MAIN, pero el paquete MAIN para GO tiene un significado especial, ya que justamente main() es la funcion principal de todo programa, por tanto GO al package main lo establecera como punto de entrada
principal de un programa ejecutable. Siempre todo programa en GO debe tener un paquete MAIN y una funcion main(), al igual que C, que es donde empieza a correr el programa
*/
package main

// Ahora vamos a importar un paquete, este paquete se llama FMT, es decir, asi como nosotros creamos el paquete MAIN,
// nosotros lo que estamos haciendo aca es importar el paquete FMT
import "fmt"

// Ahora si colocamos la funcion principal main(), para ello usamos la palabra reservada FUNC que lo que hace es declarar una funcion
func main() {
	// ¿Para que importamos el paquete FMT? Pues para claramente usar la instruccion que imprime, ya que sale del paquete FMT:
	fmt.Println("Primer programa")
	// Obviamente la funcion no retorna nada ya que no le indicamos que retorne nada
}

/*
Listo, ya codeamos nuestro primer programa!, por lo que se imprimira 'Primer programa' en STDOUT cuando compilemos y ejecutemos el codigo fuente
Ahora la gran pregunta es ¿Como se compila? Bueno para ello vamos al directorio donde esta este archivo y en la Shell se usa el compilador 'go build' en lugar de 'gcc', y a continuacion el nombre del archivo que queremos
compilar, la instruccion 'go build' lo que causa es que justamente se genere el ejecutable osea:

		>> go build primer_programa.go

Eso compilaria el codigo generando el ejecutable, luego para ejecutarlo tenemos que ingresar:

		>> ./primer_programa


Tambien si nosotros queremos compilar y ejecutar todo de manera simultanea entonces deberiamos
ingresar el comando 'go run', osea:

		>> go run primer_programa.go







Ahora vamos a ir al archivo variables.go
*/
