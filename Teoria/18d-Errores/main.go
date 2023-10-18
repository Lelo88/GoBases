package main

import (
	"os"
	"log"
)

func main(){
	// el siguiente codigo abre un archivo que coloquemos. en caso de que no lo encuentre, se lanza un error
	file, err := os.Open("not_exists.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	print(file)
}