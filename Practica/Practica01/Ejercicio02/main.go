/**/

package main

import "fmt"

func main(){

	var temperatura float64
	var humedad uint8
	var presion uint16

	temperatura = 30.6
	humedad = 75
	presion = 1084

	fmt.Println("La temperatura en el dia de hoy es de",temperatura,", con una humedad de", humedad,"y una presion de",presion)
}