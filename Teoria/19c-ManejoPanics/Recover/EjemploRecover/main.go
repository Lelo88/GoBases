package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Comienza el programa")
	function()
	fmt.Println("Todo salio bien")
}

func function() {
    var apiSlice []string
	defer func() {
		err := recover()
		if err!= nil {
            //fmt.Println("Ocurrio un error")
			log.Println("Ocurrio un error", err) //aca imprimimos el mensaje de error con fecha y hora, y el tipo de error
		}
	}()
	apiSlice = append(apiSlice, "hola") //al definir el slice, ya no va a tirar ningun error
	_=apiSlice[0]
}

