package main

import "fmt"

func main(){
	frutas :=[] string {"manzana", "banana", "pera"}

	for i, fruta := range frutas { // la primer variable que se itera es la posicion en este caso del string, el segundo es el valor en donde se ubica
		fmt.Println(i,fruta)
		if i == 1 {
			break
		}
	}
}