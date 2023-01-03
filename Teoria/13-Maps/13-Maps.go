/*Los maps son variables del tipo clave valor, definiendo un tipo de clave para las variables y uno para los valores*/

package main

import "fmt"

func main(){
	myMap := map[string]int{}
	fmt.Println(len(myMap))

	//forma de ingresar datos al map creado
	myMap["Leandro"] = 34
	myMap["Romina"] = 36
	myMap["Luna"] = 5

	fmt.Println(myMap)
	fmt.Println("Ahora el tamaño es de", len(myMap))
	fmt.Println("La edad de Leandro es de",myMap["Leandro"]) //accedo al dato del map
	
	//actualizo un dato
	
	myMap["Leandro"] = 35
	

	//recorreremos el map

	for key,value := range myMap{
		fmt.Println(key, ":", value)
	}

	delete(myMap, "Leandro")
	fmt.Println("Ahora la variable map queda de esta manera")
	for key,value := range myMap{
		fmt.Println(key, ":", value)
	}
}