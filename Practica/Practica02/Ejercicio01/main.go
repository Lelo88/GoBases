package main

import "fmt"

func main(){
	var palabra string
	fmt.Printf("Ingrese una palabra: ")
	fmt.Scanf("%s", &palabra)

	fmt.Println("")
	fmt.Println("Asi se ve la palabra letra por letra: ")
	for _,letra := range palabra{
		fmt.Printf("%c\n", letra)
	}
}