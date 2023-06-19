package main

import (
	"fmt"
)


func checkPalindrome (inputString string) bool {
	
	for i:=0; i < len(inputString); i++ {
		if inputString[i] != inputString[len(inputString) - i - 1] {
			return false
		}
	}

	return true
}

func main () {
	isPalindrome := checkPalindrome("aabaa")
	fmt.Println(isPalindrome)
	isPalindrome = checkPalindrome("abac")
	fmt.Println(isPalindrome)

	//intentar pasar una cadena con espacios y chequear si es palindromo

}