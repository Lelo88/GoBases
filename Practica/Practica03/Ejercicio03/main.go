/*
Ejercicio 3 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la 
categoría.
Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, 
la categoría y que devuelva su salario.
*/

package main

import (
	"fmt"
	"strings"
)

func ingreso() (categoria string, horasTrabajadas int, ok bool){
	fmt.Println("Ingrese la categoría del empleado")
	fmt.Scanf("%s", &categoria)
	categoria = strings.ToUpper(categoria)
	if categoria != "A" && categoria !="B" && categoria !="C" {
		return //devuelvo el status falso por mas que pase una categoria distinta
	}
	fmt.Println("Ingrese la cantidad de horas trabajadas")
	fmt.Scanf("%d", &horasTrabajadas)
	if horasTrabajadas <= 0 {
		return //devuelvo el status falso por mas que pase horas en negativo
	}
	ok = true
	
	return //devuelvo el status, el sueldo y la categoria
}

func calculaSalario(cat string, horasTrabajadas int) (salario float64) {
	switch cat {
	case "C":
		salario = float64(horasTrabajadas * 1000)
	case "B":
		salario = float64(horasTrabajadas * 1500)
		salario += salario * 0.2
	case "A":
		salario = float64(horasTrabajadas * 3000)
		salario += salario * 0.5	
	}
	return //devuelvo implicitamente salario
}

func main(){
	cat, horas, status := ingreso()

	if !status {
		fmt.Println("No puede calcularse el salario del empleado")
	} else {
		fmt.Println(cat, horas)
		fmt.Printf("El salario total correspondiente a las horas trabajadas es de %.2f.\n", calculaSalario(cat,horas))
	}
}