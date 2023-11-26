package main

import "fmt"

func main(){
	//enteros: Pueden ser con signo o sin signo. A su vez, los enteros signados pueden estar definido por un tipo especifico
	//de numeros enteros

	var entero uint = 15
	var enteroSignado int8 = -1
	
	fmt.Println(entero, enteroSignado)

	//numeros flotantes: pueden ser de 32 o 64 bits
	var flotante float64 = 23.4

	fmt.Println(flotante)

	// strings
	var cadena string = "Esta es una cadena"
	fmt.Println(cadena)

	nuevoEntero := int8(flotante)

	fmt.Println(nuevoEntero)

	numHexa := 0x11
	fmt.Println("El numero hexadecimal 0x11 es: ", numHexa) // 17
}