package main

import "fmt"

func main(){
	a:=10
	b:=2

	if a >= b {
		fmt.Println(a, "es el mayor")
	} else if a <= b {
		fmt.Println(b, "es el mayor")
	} else {
		fmt.Println("son iguales")
	}
}

