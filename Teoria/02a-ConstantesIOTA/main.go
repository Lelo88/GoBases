package main

import "fmt"

func main(){
	const (
		// puedo usar expresiones iota en las constantes para repetir las expresiones
		// de manera predeterminada, comienza en 0. iota es un contador de constantes que se incrementa en 1.
		uno = iota + 1
		dos
		tres
	)

	fmt.Println(uno, dos, tres)

	const (
		EST = -(5 + iota) // EST = -5
		_ // con esto tendría una constante inexistente que seria la siguiente a EST
		MST  // MST = 7
		PST // PST = 8
	)

	fmt.Println(EST, MST, PST)
}