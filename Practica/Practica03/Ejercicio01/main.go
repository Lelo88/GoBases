package main

import "fmt"

func calculaImpuesto(sueldo float64) (descuento float64, ok bool){
	if sueldo < 50000 {
		return 
	} else if sueldo >= 50000 && sueldo<150000 {
		ok = true
		descuento = sueldo * 0.17
		return
		} else {
		ok = true
		descuento = sueldo * 0.27 
		return
	}
}

func main(){
	sueldo :=0.0
	fmt.Printf("Ingrese el sueldo correspondiente: ")
	fmt.Scanf("%f", &sueldo)

	descuento, ok := calculaImpuesto(sueldo)

	if !ok {
		fmt.Println("No se realiza ningun descuento sobre el sueldo ingresado")
	} else {
		fmt.Printf("Al sueldo ingresado le corresponde un descuento de %.2f\n", descuento)
	}

}