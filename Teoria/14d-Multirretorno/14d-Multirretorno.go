//en las funciones tenemos dos pares de parentesis: el primero es el que recibe los parametros. el segundo es el que devuelve

package main

import (
	"errors"
	"fmt"
)

func suma(num1, num2 float64) float64{
	return num1 + num2
}

func resta(num1, num2 float64) float64{
	return num1- num2
}

func sumarYrestar(num1,num2 float64) (float64, float64){ //en go puedo realizar funciones que devuelvan dos valores
	resultadoSuma := suma(num1,num2)
	resultadoResta := resta(num1, num2)

	return resultadoSuma, resultadoResta
}

//crearemos una funcion con un error a proposito para ver mejor el multirretorno
func dividir(num1, num2 float64) (float64, error){
	if num2 == 0{
		return 0, errors.New("no es posible dividir en 0")
	}

	return num1/num2, nil
}

//con esta funcion puedo retornar variables y errores a la vez declarandolos en la variable de vuelta
func dividir2(num1, num2 float64) (resultado float64, ok bool){
	
	if num2 == 0 {
		return 
	}

	ok=true
	resultado = num1/num2
	return 
}

func main(){
	fmt.Println(sumarYrestar(3,1))
	//tambien se puede llamar a la funcion de la siguiente manera
	resultadoSuma, resultadoResta := sumarYrestar(5,4) //la primer variable va a representar al primer retorno, la segunda al segundo
	fmt.Println("Resultado Suma:", resultadoSuma, "Resultado Resta:", resultadoResta) 
	//se pueden ignorar valores de retorno
	_,otraResta := sumarYrestar(8,9)
	fmt.Println("Ignoro el resultado de la suma:", otraResta)
	
	fmt.Printf("\n")
	a,b := 6.0,0.0
	resultadoDiv, err := dividir(a,b)
	
	if err != nil {
		fmt.Println(err) //si el error es distinto de nil, entonces me devolvera el error en la division
	} else {
		fmt.Println(resultadoDiv) //me imprime el resultado en caso de que la division se efectue correctamente
	}

	otroResultDiv, ok := dividir2(a,b)

	if !ok {
		panic("Ha ocurrido un error")
	}

	fmt.Println(otroResultDiv)
}