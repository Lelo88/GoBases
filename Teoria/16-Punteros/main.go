//punteros: variable a la cual podemos acceder a su posicion de memoria y a su valor tambien, asi como tambien modificarlo
//a traves de funciones o metodos y que este quede con el valor actualizado

package main

import "fmt"

func incrementar(v *int){ //esta funcion nos permitira incrementar en uno el valor v. dado que las funciones retornan variable,
	*v++				  //con los punteros podemos acceder al valo de la variable en memoria y modificarla. 
}

func main(){
//1. tenemos la siguiente variable
	v:= 19

	fmt.Println("v vale:", v)

	incrementar(&v) //llamamos a la funcion v pasando como argumento la variable con &, es decir, su posicion en memoria para que pueda ser modificada

	fmt.Println("Ahora v vale: ", v)
}