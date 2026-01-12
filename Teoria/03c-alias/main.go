package main

import "fmt"

func main() {
	// Los alias son tipos que se crean a partir de un tipo existente
	// A diferencia de los tipos definidos, los alias no crean un nuevo tipo

	var a uint8 = 30 // uint8 es un alias de uint8
	var b byte       // byte es un alias de uint8

	b = a 
	_ = b
	
	type second = uint
	var hour second = 3600
	fmt.Println("Minutes in an hour:", hour/60) // 
}

