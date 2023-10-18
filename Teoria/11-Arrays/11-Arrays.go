/*Los arrays son tipos de datos que nos permiten almacenar un conjunto de datos homogeneo, todos del mismo tipo*/

package main

import "fmt"

func main(){
	var a[2] string //al momento que creamos el array tenemos que definirle el tamaño
	
	//voy a agregarle unos valores de ejemplo

	for i:=0;i<2;i++{
		fmt.Println("Ingrese la cadena", i+1)
		fmt.Scanf("%s", &a[i])
	}

	fmt.Printf("\n")

	for i:=0;i<2;i++{
		fmt.Println(a[i])
	}

	/*Si quiero acceder a un espacio de memoria que no esta definido, go enseguida me remarcará el error*/
}