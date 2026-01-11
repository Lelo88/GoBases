package main

import "fmt"

//funciones anonimas

func main(){

	myfunc:= func (){
		fmt.Println("Hola desde una funcion anonima")
	}

	myfunc()
}

// las funciones anonimas son utiles cuando se necesita crear una funcion que se va a usar una sola vez
// y no se quiere definir una funcion con nombre
// tambien son utiles cuando se necesita pasar una funcion como argumento a otra funcion
// por ejemplo, en el paquete "time" se usa una funcion anonima para crear un timer
// tambien se usan en los metodos de los structs
