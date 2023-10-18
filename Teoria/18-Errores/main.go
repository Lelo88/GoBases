//Los errores en Go no se manejan como en otros lenguajes de programacion, sino que se manejan como interfaces

package main

import (
    "fmt"
    "errors"
)

func validateStateCoded(code int) error {
	if code > 399 {
		return errors.New("unexpected http status code") //usamos el paquete errors con su funcion new para lanzar el error en caso de que no se cumpla una condicional
	}
	return nil
}

func main(){

	err := validateStateCoded(404) 
	if err!= nil{
		fmt.Printf("http request failed: %v\n", err)
		return
	}
	
	fmt.Println("the program ended successfully")
}