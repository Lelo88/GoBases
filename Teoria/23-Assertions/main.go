/*Una aserción de tipo proporciona acceso al valor concreto subyacente de un valor de interfaz.

t := i.(T)
Esta declaración afirma que el valor de interfaz i contiene el tipo concreto T y 
asigna el valor T subyacente a la variable t.

Si i no tiene una T, la declaración provocará pánico.

Para probar si un valor de interfaz contiene un tipo específico, 
una aserción de tipo puede devolver dos valores: el valor subyacente 
y un valor booleano que informa si la aserción tuvo éxito.

t, ok := i.(T)
Si tengo una T, entonces t será el valor subyacente y ok será verdadero.

Si no, ok será falso y t será el valor cero del tipo T, y no se producirá pánico.

Tenga en cuenta la similitud entre esta sintaxis y la de leer de un mapa.*/

package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*f = i.(float64) // panic
	fmt.Println(f)*/
}

