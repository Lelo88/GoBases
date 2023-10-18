package main

import "fmt"

func main(){
	fmt.Println("La suma de 4 y 5 es de",suma(4,5)) //una manera de llamar a la funcion es de manera directa 

	resultado := suma(6,7) //otra manera de llamar a una funcion, asociandola a una variable
	fmt.Println("El resultado entre 6 y 7 es", resultado)
} 

func suma(a, b int) int { //nombre de la funcion (parametros a pasar (maximos 3 es lo recomendado)) retorno de la funcion {cuerpo}
	return a + b
}

