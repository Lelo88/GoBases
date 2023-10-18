package main

import "fmt"

func main(){
	option := 0
	fmt.Println("Ingrese una opcion: ")
	fmt.Scanf("%d", &option)

	switch option {
		case 1:
		fmt.Println("Selecciono el uno")
		case 2: 
		fmt.Println("Selecciono el dos")
		case 3:
		fmt.Println("Selecciono el tres")
		default:
		fmt.Println("Ninguno de los números")
	}
}