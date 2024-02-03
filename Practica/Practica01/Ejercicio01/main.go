/*
1. Crear una aplicación donde tengas como variable tu nombre y dirección.
2. Imprimir en consola el valor de cada una de las variables.

*/

package main

import "fmt"

func main() {

	var (
		nombre    = "Leandro"
		direccion = "Calle Falsa 123"
	)

	fmt.Println("Mi nombre es", nombre, "y vivo en", direccion)
}
