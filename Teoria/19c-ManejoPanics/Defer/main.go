package main

import "fmt"

func main() {
	defer func(){
		fmt.Println("Esta funcion se ejecuta a pesar de producirse un panic")
	}()
	
	panic("se produjo un panic")

}

