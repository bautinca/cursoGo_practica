package main

import (
	"fmt"
	"os"
)

// Vamos a implementar una funcion para leer un archivo asi aprendemos cosas nuevas de GO

func main() {

	// Primero abrimos el archivo con la funcion Open() del paquete OS
	f, err := os.Open("texto_defer2.txt")

	// Como vemos, la funcion Open() nos da 2 valores, el puntero al archivo abierto y si hay error o no
	// en caso de que no haya error entonces err = nil
	// Entonces obviamente tenemos que verificar si al abrir al archivo dio error
	if err != nil {
		panic(err) // Luego veremos que hace, pero basicamente detiene la ejecucion del programa
	}

	// Lo bueno del DEFER es que por ejemplo puedo cerrar el archivo en esta linea ya que nos aseguramos
	// que se ejecute justo al finalizar el main el close del archivo
	defer f.Close() // Entonces nos aseguramos que SI O SI se ejecute el defer Close ya sea si el main llega al final o si en el camino encuentra un error
	// ya que los defer tambien se ejecutan cuando encontramos un error

	// Si no ocurrio ningun error entonces seguimos, vamos a crear un slice de tipos de datos bytes que va a tener una capacidad de 175
	data := make([]byte, 175)
	// Ahora si abrimos el archivo logicamente que sigue? LEERLO! Para ello usaremos la funcion READ() donde
	// toda la info. que se lea se volcara en el slice 'data'
	// Cabe aclarar que la funcion READ() retorna 2 datos, un entero que indica la cantidad de bytes leidos, y si dio error o no
	c, err := f.Read(data)

	// Verificamos si en la lectura dio error o no
	if err != nil {
		panic(err) // Nuevamente usamos panic() luego veremos que hace
	}

	// Ahora finalmente imprimimosel contenido leido
	fmt.Printf("Cantidad de bytes leidos: %d\nTexto leido:\n%s\nerror: %v", c, data, err)
}

// Ahora vamos a ir al archivo panic_recover.go
