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

	//crear un map con make
	map2 := make(map[string]int)
	map2["Juan"] = 25
	map2["Maria"] = 30
	fmt.Println("Mapa creado con make:", map2)

	//crear un map con make y definir el tamaño
	map3 := make(map[string]int, 10)
	map3["Juan"] = 25
	map3["Maria"] = 30
	fmt.Println("Mapa creado con make y definir el tamaño:", map3)

	// La diferencia entre make(map[string]int) y make(map[string]int, 10) es que el segundo define un tamaño inicial para el map
	// Esto puede ser útil para optimizar el rendimiento cuando se sabe que el map tendrá un número determinado de elementos
}