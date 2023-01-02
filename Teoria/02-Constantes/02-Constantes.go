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
}