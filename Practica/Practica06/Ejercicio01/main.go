package main

import "fmt"

type Error struct{
}

func (e Error) Error() string{
	return "Error: el salario ingresado no alcanza el mínimo no imponible"
}

func main(){
	salary:=0
	fmt.Println("Ingrese el salario correspondiente: ")
	fmt.Scanf("%d", &salary)
	
	if salary < 150000 {
		fmt.Println(Error{})
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}

