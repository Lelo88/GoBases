package main

import "fmt"

//funciones anonimas

func main(){

	myfunc:= func (){
		fmt.Println("Hola desde una funcion anonima")
	}

	myfunc()
}