package main

import (
	"fmt"
	"errors"
)

type MyError struct{}

func (e MyError) Error() string {
    return "Error: el salario menor a 10000"
}

func validaSalario(salario int)	error {
	if salario <= 10000 {
		return MyError{}
	}
	return nil
}

func main() {
	mySalary := 0
	
	fmt.Println("Ingrese el valor de su salario")
	fmt.Scanf("%d", &mySalary)

	err:= validaSalario(mySalary)

	if errors.Is(err, MyError{}) {
		fmt.Println(err)
	}
}
