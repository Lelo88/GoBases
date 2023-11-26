package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	name := "Inanç" 
	fmt.Println(len(name)) // la longitud es de 6 a pesar de que sean 5 caracteres por el ultimo caracter
	fmt.Println(utf8.RuneCountInString(name)) // cuenta los caracteres no convencionales y los convencionales, dando la longitud exacta
}

