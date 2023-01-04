package main

import "fmt"

const (
	OperadorSuma  = "+"
	OperadorResta = "-"
	OperadorMult  = "*"
	OperadorDiv   = "/"
)


func sumar(nums ... float64) (resultado float64) {
	for index := range nums{
		resultado += nums[index]
	}
	return
}

func restar(nums ... float64) (resultado float64) {
	for index := range nums{
		resultado -= nums[index]
	}
	return
}

func multiplicar(nums ... float64) (resultado float64) {
	for index := range nums{
		resultado *= nums[index]
	}
	return
}

func dividir(nums ... float64) (resultado float64) {
	for index := range nums{
		if nums[index] == 0 {
			return 
		}
		resultado /= nums[index]
	}
	return
}

func operacionAritmetica(operador string, nums ...float64) (resultado float64){
	switch operador{
	case OperadorSuma:
		resultado = sumar(nums ...)
	}
	return resultado
}

func main(){
	fmt.Println()
}