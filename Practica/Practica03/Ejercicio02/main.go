/*Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. 
No se pueden introducir notas negativas.*/

package main

import "fmt"

func calculaPromedio(notas ...int) (float64){
	sumaNot := 0

	for _,value := range notas {
		sumaNot += value
	}

	return float64(sumaNot)/float64(len(notas))
}

func ingreso() (status bool, notas[] int){
	
	fmt.Println("Ingrese las notas correspondientes al alumnado: ")
	for {
		nota := 0
		fmt.Scanf("%d", &nota)
		if nota < 0 || nota > 10 {
			break
		}
		notas = append(notas, nota)
	} 

	if len(notas) == 0 { 
		return
	}
	status = true
	return 
}

func main(){
	
	ok,notas := ingreso()

	if !ok {
		panic("No se puede realizar el calculo")
	}
	
	fmt.Println(notas)

	fmt.Printf("El promedio de las notas ingresadas es de %.2f\n", calculaPromedio(notas...))
}

//se puede optimizar el codigo para realizar el ingreso por funcion y llamarlo en el main