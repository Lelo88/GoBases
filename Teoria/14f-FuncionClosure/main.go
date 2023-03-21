package main

import (
	"fmt"
	"strings"
)
//una funcion closure tiene una funciona anidada que devuelve un valor especificado en la funcion padre
func repeat (n int) func(cadena string) string{
	return func(cadena string) string{
		return strings.Repeat(cadena, n)
	}
} 

func main() {
	repeat3 := repeat(3)
	fmt.Println(repeat3("hola"))
	fmt.Println(repeat3("mundo"))
}