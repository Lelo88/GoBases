package main

import "fmt"

func main(){
	//maneras de definir una variable
	
	var nombre string = ""
	//asignando un valor a una variable
	nombre = "Leandro"
	fmt.Println(nombre)

	//declaracion de multiples variables

	producto, precio := "remera", 20
	fmt.Println(producto, precio)

	//tambien se pueden declarar de esta manera

	var (
		cantidad = 35
		otroPrecio = 24.4
		enStock = true
	)

	fmt.Println(cantidad, otroPrecio, enStock)

	//declaracion corta de una variable
	//internamente, go reconoce el tipo de la variable al definirla de esta manera
	horas := 30
	fmt.Println(horas)
}