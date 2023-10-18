package main

import (
	"fmt"
	"errors"
)

const (
	Dog 	= "perro"
	Cat 	= "gato"
	Spider 	= "araña"
	Hamster = "hamster"
)

type Animal func()

func alimentoPerro(cantidad float64) float64{
	return cantidad * 10
}

func alimentoGato(cantidad float64) float64 {
	return cantidad * 5
}

func alimentoHamster (cantidad float64) float64 {
	return cantidad * 0.25
}

func alimentoTarantula (cantidad float64) float64{
	return cantidad * 0.15
}

func animal(opcion string) (func(float64) float64, error){
	switch  opcion {
		case Dog:
			return alimentoPerro, nil
		case Cat:
			return alimentoGato, nil
		case Hamster: 
			return alimentoHamster, nil 
		case Spider:
			return alimentoTarantula, nil
		default:
			return nil, errors.New("Animal incorrecto")
	}
} 

func main(){
	perro, msg := animal(Dog)
	gato, msg1 := animal(Cat)
	araña, msg2 := animal(Spider)
	hamster, msg3 := animal(Hamster)
	//paloma, msg4 := animal("paloma")

	if msg!= nil && msg1!= nil && msg2!= nil && msg3 != nil{
		fmt.Println(msg)
	} 

	cantidad := perro(19)
	cantidad += gato(20)
	cantidad += araña(10)
	cantidad += hamster(25)
	

	fmt.Println(cantidad)
}