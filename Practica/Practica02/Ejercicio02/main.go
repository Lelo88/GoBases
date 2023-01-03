package main

import "fmt"

func main(){
	nombre, sueldo, antiguedad, edad := "", 0, 0,0
	var esEmpleado string
	fmt.Println("Ingrese el nombre del empleado: ")
	fmt.Scanf("%s", &nombre)

	fmt.Println("Ingrese su edad:")
	fmt.Scanf("%d", &edad)

	fmt.Println("Ingrese la antiguedad del empleado en la empresa:")
	fmt.Scanf("%d", &antiguedad)

	fmt.Println("Ingrese el sueldo del empleado")
	fmt.Scanf("%d", &sueldo)

	fmt.Println("Se encuentra empleado?(s/n)")
	fmt.Scanf("%s", &esEmpleado)

	if edad >=22{
		if esEmpleado == "S" {
			if antiguedad > 1{
				if sueldo > 100000{
					fmt.Println("Al empleado", nombre,"se le otorgará un prestamo sin intereses")
				} else {
					fmt.Println("Al empleado", nombre,"se le otorgará un prestamo con intereses")
				}
			} else {
				fmt.Println("No puede darle prestamo")
			}
		}else {
			fmt.Println("No puede darle prestamo")
	 	}
	} else {
		fmt.Println("No puede darle prestamo")
	}
}