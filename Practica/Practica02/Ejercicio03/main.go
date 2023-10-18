package main

import "fmt"

func main(){
	var mes_ingresado int

	var meses = map[int]string{1:"Enero", 2:"Febrero", 3:"Marzo", 4:"Abril", 5:"Mayo", 6:"Junio", 
								7:"Julio",8:"Agosto",9:"Septiembre",10:"Octubre", 11:"Noviembre", 12:"Diciembre"}


	fmt.Println("Ingrese el mes que desea buscar")
	fmt.Scanf("%d", &mes_ingresado)

	for valor, mes := range meses {
		if mes_ingresado == valor {
			fmt.Println("Nº de mes ingresado:", mes_ingresado,"Mes:",mes)
			break
		} 
	}	
	
	if mes_ingresado < 1 || mes_ingresado > 12 {
		fmt.Println("El mes ingresado no corresponde a ningun mes")
	}
}