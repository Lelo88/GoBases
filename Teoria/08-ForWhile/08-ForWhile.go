package main

import "fmt"

func main(){
	sum := 1
	
	fmt.Println("Sum vale en este momento: ", sum)
	
	for sum <=10 {
		sum+=sum
	}
	fmt.Println("El valor de sum ahora es de:", sum)

	
}