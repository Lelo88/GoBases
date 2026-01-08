package main

import (
	"fmt"
	"math"
)

func main() {
	var i int8 = 127
	fmt.Println("int8 overflow:", i+1) // overflow: -128
    i++
    fmt.Println("int8 overflow:", i) // overflow: -127
	i--
	fmt.Println("int8 overflow:", i) // overflow: -128
	
	var j uint8 = 255
	fmt.Println("uint8 overflow:", j+1) // overflow: 0
	j++
	fmt.Println("uint8 overflow:", j) // overflow: 0
	j--
	fmt.Println("uint8 overflow:", j) // overflow: 255

	// constant math.MaxFloat32 is the largest finite value representable by float32
	// Como es una constante, no hay overflow en la asignación
	k := float32(math.MaxFloat32)
	fmt.Println("float32:", k)
	fmt.Println("float32 overflow:", k+1)
}
