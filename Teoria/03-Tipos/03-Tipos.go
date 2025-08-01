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

	// numeros complejos: se pueden definir numeros complejos de 64 bits
	var complejo complex64 = 1 + 2i // 1 es la parte real y 2 es la parte imaginaria
	fmt.Println("El numero complejo es: ", complejo)

	// se pueden hacer operaciones con numeros complejos
	real := real(complejo) // obtiene la parte real del numero complejo
	imag := imag(complejo) // obtiene la parte imaginaria del numero complejo
	fmt.Println("La parte real es: ", real)
	fmt.Println("La parte imaginaria es: ", imag)

	// se pueden hacer operaciones con numeros complejos
	suma := complejo + 2 + 3i // suma de numeros complejos
	fmt.Println("La suma de los numeros complejos es: ", suma)

	// se pueden hacer operaciones con numeros complejos
	multiplicacion := complejo * (2 + 3i) // multiplicacion de numeros complejos
	fmt.Println("La multiplicacion de los numeros complejos es: ", multiplicacion)

	// se pueden hacer operaciones con numeros complejos
	division := complejo / (2 + 3i) // division de numeros complejos
	fmt.Println("La division de los numeros complejos es: ", division)

	// se pueden hacer operaciones con numeros complejos
	resta := complejo - (2 + 3i) // resta de numeros complejos
	fmt.Println("La resta de los numeros complejos es: ", resta)	


	// runes: son un tipo de dato que representa un caracter unicode
	var runeChar rune = 'A' // se puede definir un caracter unicode con comillas simples
	fmt.Println("El caracter unicode es: ", runeChar)
	fmt.Println("El caracter unicode es: ", string(runeChar)) // se puede convertir a

	// bytes: son un tipo de dato que representa un byte
	var byteChar byte = 'B' // se puede definir un byte con comillas simples
	fmt.Println("El byte es: ", byteChar)
	fmt.Println("El byte es: ", string(byteChar)) // se puede convertir a string
}