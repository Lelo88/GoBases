/*los slices son como los arrays, constituidos por datos homogeneos. La diferencia radica en que su tamaño puede variar*/

package main

import "fmt"

func main(){
	var slice = []bool{true, false} //primer manera de definir un slice
	sliceInt := make([]int, 5) //otra manera de definir slices con la funcion make, para indicar el tipo y el tamaño
	sliceInt2 := []int{1,2,3,4,5}

	for _,value := range slice {
		fmt.Println(value)
	}

	for i:=0; i<5 ; i++ {
		fmt.Println("Ingrese numeros")
		fmt.Scanf("%d", &sliceInt[i])
	}

	for _, value := range sliceInt{
		fmt.Println(value) // value es el valor del slice
	}

	fmt.Println(sliceInt[1:4]) // sliceInt[1:4] es un slice que contiene los elementos desde el índice 1 hasta el índice 3

	fmt.Println("Nuevo slice: ", sliceInt2)
	fmt.Println("Tamaño de sliceInt2: ", len(sliceInt2), "capacidad: ", cap(sliceInt2))
	sliceInt2 = append(sliceInt2, 6,7,8) // append agrega elementos al slice
	fmt.Println("Nuevo valor del nuevo slice: ", sliceInt2) // sliceInt2 ahora tiene 8 elementos
	fmt.Println("Tamaño de sliceInt2: ", len(sliceInt2), "capacidad: ", cap(sliceInt2)) //la capacidad aumenta por realocacion
	
	slice2 := append(sliceInt2, 9, 10, 11)
	fmt.Println("Nuevo slice: ", slice2)
	fmt.Println("Tamaño de slice2: ", len(slice2), "capacidad: ", cap(slice2))
	// La capacidad del slice2 es 10, que es el doble de la capacidad del sliceInt2
	// Esto se debe a que Go duplica la capacidad cuando se supera el límite
	// Al agregar más elementos, Go duplica la capacidad para evitar realocaciones frecuentes. 
}
