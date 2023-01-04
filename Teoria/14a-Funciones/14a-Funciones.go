//Ejemplo de calculadora simple en funciones pasando diversos parametros

package main

import "fmt"

const (
	Suma  = "+"
	Resta = "-"
	Mult  = "*"
	Div   = "/"
)

func main(){
	/*Hago distintos llamados a la funcion, pasando como parametros dos numeros y un operador declarado en la constante de arriba*/
	fmt.Println(operacionAritmetica(4,5,Suma))
	fmt.Println(operacionAritmetica(4,5,Resta))
	fmt.Println(operacionAritmetica(4,5,Mult))
	fmt.Println(operacionAritmetica(4,5,Div))
}

func operacionAritmetica (valor1, valor2 float64, operador string) float64{ //pasamos dos numeros y un string que se referira al operador
	switch operador{
	case Suma:
		return valor1 + valor2
	case Resta:
		return valor1 - valor2
	case Mult:
		return valor1 * valor2
	case Div:
		if valor2 !=0 {
			return valor1 / valor2
		}
	}
	return 0
}