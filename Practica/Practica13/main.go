package main

import "fmt"

// Give an array of integers, find the pair of adjacent elements that has the largest product and return
// that product.

func productNumbersAdjacent (inputArray []int) int {
	
	if len(inputArray) < 2 {
		return 0
	}
	
	var product int = 0
	var aux int = 0
	
	for i:=0; i < len(inputArray) - 1; i++ {
		aux = inputArray[i] * inputArray[i+1]
		if aux > product {
			product = aux
		}
	}
	
	return product
}

func main () {
	
	var inputArray []int = []int{3, 6, -2, -5, 7, 3}
	
	var product int = productNumbersAdjacent(inputArray)
	
	fmt.Println(product)

}