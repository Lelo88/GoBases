/*Un cambio de tipo es una construcción que permite varias aserciones de tipo en serie.

Un cambio de tipo es como una declaración de cambio normal, pero los casos en un cambio de 
tipo especifican tipos (no valores), y esos valores se comparan con el tipo del valor que 
tiene el valor de interfaz dado.

La declaración en un cambio de tipo tiene la misma sintaxis que una aserción de tipo i.(T), 
pero el tipo T específico se reemplaza con la palabra clave type.

Esta declaración de cambio comprueba si el valor de interfaz i tiene un valor de tipo T o S. 
En cada uno de los casos T y S, la variable v será de tipo T o S respectivamente y tendrá el valor que tiene i. 
En el caso predeterminado (donde no hay coincidencia), la variable v es del mismo tipo de interfaz 
y valor que i.*/

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}