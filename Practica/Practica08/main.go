package main

import "fmt"

func compareTriplets(a,b []int32) []int32{

	var result = make([]int32, 2)
	//var sum int32 = 0

	for i:= 0; i < 3; i++ {
		switch{
			case a[i] < b[i]:
				result[0] += 1
			case a[i] > b[i]:
				result[1] += 1
		}
	}
	return result
}

func main() {
	var a = []int32 {1,1,3}
	var b = []int32 {4,2,1}

	fmt.Println(compareTriplets(a,b))
}