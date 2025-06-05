package main

import "fmt"

func main(){
	//declaracion de una constante
	const status = "ok"

	//bloque de constante: se declara igual que una variable

	const (
		producto = "remera"
		precio = 25.5
		cantidad = 10
	)
	
	fmt.Println(status, producto, precio, cantidad)

	// constante no definida
	 const sinDefinir = 10
	 fmt.Println(sinDefinir)

	// constante definida
	 const definido int64 = 10
	 fmt.Println(definido)

	 // Si la constante no tiene tipo, el compilador lo infiere. Esto quiere decir que el compilador le asigna un tipo a la constante.
	 // y si no se le asigna un tipo, el compilador le asigna el tipo de la variable que se le asigna.
}