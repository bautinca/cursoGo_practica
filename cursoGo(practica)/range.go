package main

import "fmt"

func main() {
	// Nos acordamos de RANGE? En Python lo teniamos, y casi siempre iba con el FOR!
	// Osea for ... in range(...)

	// Bueno en GO volvemos a tener el RANGE!

	// Primero lo que vamos a hacer es crearnos un slice de strings
	nombres := []string{"Alejandro", "Maria", "Pedro", "Carlos"}

	// Ahora, como hacemos para recorrer la variable nombres? Lo podemos hacer con range! de una manera elegante

	for index, nombre := range nombres {
		fmt.Printf("El indice es %d y el nombre es %s\n", index, nombre)
	}

	fmt.Println("-----------------------------------------")

	// Solo por curiosidad tambien lo podriamos haber recorrido asi pero es un poco mas desprolijo:
	for i := 0; i < len(nombres); i++ {
		fmt.Printf("El indice es %d y el nombre es %s\n", i, nombres[i])
	}

	// En fin, el range nos brinda como una manera mas 'prolija' y clara para codear
	fmt.Println("-----------------------------------------")

	// Ahora, si o si tengo que tomar el indice? No!, si quiero puedo solamente tomar los nombres:
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	// Pero CUIDADO! Esto no se puede hacer para tomar solo los indices, es decir, o se puede hacer indice,_ eso esta mal
	// Lo que si podemos hacer es colocar directamente el indice, osea:
	fmt.Println("-----------------------------------------")

	for indice := range nombres {
		fmt.Println(indice)
	}

	fmt.Println("-----------------------------------------")

	// Ahora tambien podemos simplemente recorrer strings! Veamos:
	cadena_formateada := fmt.Sprintf("Esta es una cadena codigo %d formateada, y mi nombre es %s", len(nombres), nombres[0])

	for _, caracter := range cadena_formateada {
		fmt.Printf("%c\n", caracter)
	}

}

// Ahora vamos al archivo maps.go
