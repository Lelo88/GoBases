package main

import (
	"fmt"
	"errors"
)

var errSalarioMinimo = errors.New("salario menor a 10000")

func validaSalario(salario int)	error {
	if salario <= 10000 {
		return errSalarioMinimo
	}
	return nil
}

func main() {
	mySalary := 0
	
	fmt.Println("Ingrese el valor de su salario")
	fmt.Scanf("%d", &mySalary)

	err:= validaSalario(mySalary)

	if errors.Is(err, errSalarioMinimo) {
		fmt.Println(err)
	}
}
