package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/*n valor de interfaz nulo no contiene ni valor ni tipo concreto.

Llamar a un método en una interfaz nula es un error de tiempo de 
ejecución porque no hay ningún tipo dentro de la tupla de la interfaz para indicar qué método concreto llamar.*/