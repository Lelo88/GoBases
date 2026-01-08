package main

import (
	"fmt"
	"strconv"
)

func main() {

	// entero
	string_number := fmt.Sprintf("%d", 65)
	fmt.Println(string_number)	

	// flotante
	string_number = fmt.Sprintf("%f", 65.5)
	fmt.Println(string_number)

	// binario
	string_number = fmt.Sprintf("%b", 65)
	fmt.Println(string_number)

	// hexadecimal
	string_number = fmt.Sprintf("%x", 65)
	fmt.Println(string_number)

	// complejo
	string_number = fmt.Sprintf("%c", 65)
	fmt.Println(string_number)

	// ahora con parse
	number, err := strconv.ParseInt("10", 10 ,64) // 10 es el sistema de numeracion, 10 es el numero de bits y 64 es el numero de bits 
	fmt.Println(number, err)

	// ahora con atoi
	number_atoi, err := strconv.Atoi("10")
	fmt.Printf("number_atoi type is %T, value %v\n", number_atoi, number_atoi)

	// ahora con itoa
	string_number_2 := strconv.Itoa(20)
	fmt.Printf("string_number_2 type is %T, value %q\n", string_number_2, string_number_2)
}
