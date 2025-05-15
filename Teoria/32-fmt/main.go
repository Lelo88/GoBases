package main

import "fmt"

func main() {

	name := "Leandro"
	fmt.Println("Hola " + name)

	age := 36
	fmt.Println("Tengo " + fmt.Sprint(age) + " años")

	fmt.Printf("Hola %s, tengo %d años\n", name, age)

	// Imprimo el tipo
	fmt.Printf("El tipo de la variable name es: %T\n", name)
}