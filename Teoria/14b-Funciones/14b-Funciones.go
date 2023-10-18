//funciones con elipsis
//los elipsis tienen que estar al final de la declaracion de parametros

package main

import "fmt"

func sumaConElipsis(nums ...float64) (resultado float64){
	for i:= range nums {
		resultado+=nums[i]
	}
	return resultado
}

//2 ejemplo de paso de paramentros con una variable y un slice
//implementamos una funcion saludar

func saludar(name string){
	fmt.Println("Hola", name)
}

//ahora implementamos una funcion que, al pasar parametros, podamos usar la funcion saludar y la de suma con elipsis
func saludarYsumar(name string, nums... float64){
	saludar(name)
	fmt.Println(sumaConElipsis(nums...))
}

func main(){
	//fmt.Println(sumaConElipsis(2,3,4,5))
	nums := []float64{1,2,3,4,5}
	fmt.Println(sumaConElipsis(nums...))
	fmt.Println("Aca utilizo la funcion saludarYsumar")
	saludarYsumar("Leandro", nums...) 
}