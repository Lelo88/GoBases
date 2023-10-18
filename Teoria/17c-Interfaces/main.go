package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


/*Llamar a un método en un valor de interfaz ejecuta el método del mismo nombre en su tipo subyacente.

Si el valor concreto dentro de la propia interfaz es nulo, el método se llamará con un receptor nulo.

En algunos idiomas, esto desencadenaría una excepción de puntero nulo, pero en Go es común escribir 
métodos que manejen correctamente las llamadas con un receptor nulo (como con el método M en este ejemplo).*/