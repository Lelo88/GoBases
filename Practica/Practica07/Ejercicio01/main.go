package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err:= recover(); err!=nil{
			
			fmt.Println("Tipo de error: ", err)
		}
		fmt.Println("Finalizando ejecución...")
	}()

	_,err:= os.ReadFile("customers.txt")
	if err != nil {
		panic("No se encuentra el archivo")
	}
}

