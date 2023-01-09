package main

import "fmt"

func main() {
	defer func(){
		fmt.Println("Soy la funcion diferida 1")
	}()
	function() //La ultima funcion que entra es la primera que se llama. La primera es la ultima
	
	fmt.Println("Todo salio bien")

}

func function() {
	defer func(){
		//fmt.Println("Soy la funcion diferida 2")
		panic("Ocurrio un error fatal")
	}()
}

