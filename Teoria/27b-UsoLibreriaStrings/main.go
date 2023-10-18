package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	// vamos a realizar este ejemplo utilizando la entrada por argumentos
	// la misma se hace con go run main.go argumento1 argumento2 argumento3 argumento N

	messageFromArgument := os.Args[1]

	lengthMessage := len(messageFromArgument)
	repeatExclamationForMessage := strings.Repeat("!", lengthMessage)
	messageUpper := strings.ToUpper(repeatExclamationForMessage)

	fmt.Println(messageUpper)
}