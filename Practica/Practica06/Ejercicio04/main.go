package main

import (
	"fmt"
)

func validaSalario(salario int)	error {
	if salario < 150000 {
		return fmt.Errorf("el minimo no imponible es de 150000 y el salario es de %d", salario)
	}
	return nil
}

func main() {
	salario:= 0
	fmt.Println("Ingrese su salario")
	fmt.Scanf("%d", &salario)

	err:=validaSalario(salario)

	if err!=nil{
		fmt.Println(err)
	}
}
