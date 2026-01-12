package main

import "fmt"

func main() {
	type age int // age es un tipo nuevo basado en int

	// esto se utiliza para crear tipos personalizados. 
	var edad age = 30 // edad es de tipo age
	fmt.Println(edad)

	var oldAge age = 25 // oldAge es de tipo age
	fmt.Println(oldAge)

	var otraEdad int = 30
	fmt.Println(otraEdad)
	
	// edad = otraEdad // error de compilación
	edad = age(otraEdad) // conversión explícita
	fmt.Println(edad)

	// otraEdad = edad // error de compilación
	otraEdad = int(edad) // conversión explícita
	fmt.Println(otraEdad)
}

