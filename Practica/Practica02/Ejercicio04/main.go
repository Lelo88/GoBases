package main

import "fmt"

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Edad de Benjamin:",employees["Benjamin"])

	fmt.Println("Estos son los empleados mayores a 21: ")

	for empleado,edad := range employees{
		if edad > 21 {
			fmt.Println("Empleado: ", empleado, "Edad:", edad)
		}
	}

	employees["Federico"] = 25

	fmt.Println("Nuevo listado de empleados:", employees)

	delete(employees,"Pedro")

	fmt.Println("Nuevo listado de empleados despues de Pedro:", employees)
}