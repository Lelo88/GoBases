package main

import "fmt"

func main() {
	suma := 0

	for i:=1; i <= 1000 ; i++ {
		if i%3 == 0 || i%5==0 {
			suma += i
		}
	}

	fmt.Println("La suma de los multiplos de 3 y 5 menores a mil es ", suma)
}