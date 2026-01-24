package main

import "fmt"

func main() {
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	
	circumference := 2 * pi * radius
	
	fmt.Printf("Shape: %q\n", shape) // si no se pasa el valor, no se imprime, arrojando un missing en la terminal
	fmt.Printf("Circle circumference with radius %.2f: %.2f\n", radius, circumference)
	//_ = shape
}

