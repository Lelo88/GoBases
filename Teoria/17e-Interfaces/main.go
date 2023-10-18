package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*El tipo de interfaz que especifica cero métodos se conoce como interfaz vacía:

interfaz{}
Una interfaz vacía puede contener valores de cualquier tipo. (Cada tipo implementa al menos cero métodos).

Las interfaces vacías son utilizadas por código que maneja valores de tipo desconocido. 
Por ejemplo, fmt.Print toma cualquier número de argumentos de tipo interfaz{}.*/