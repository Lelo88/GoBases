package main

import "fmt"

type Persona struct { //definicion de un struct
	Nombre 	  string
	Genero 	  string
	Edad   	  int
	Profesion string
	Gustos	  Preferencias //ubicamos aca la estructura que se encuentra debajo 
}

//4. tambien podemos declarar una estructura que puede ser un campo de otra estructura

type Preferencias struct {
	Comidas		string
	Peliculas 	string
	Series		string
	Animes 		string
	Deportes 	string
}

func main(){
	//1. instancia de un objeto, en este caso de tipo Persona

	pers := Persona {"Leandro","Masculino",34,"Software Developer", Preferencias{"Tortilla de Papas", "", "", "", ""}}

	fmt.Println(pers.Edad) //2. podemos acceder a un campo en especifico del objeto instanciado
	fmt.Println(pers)
	fmt.Println(pers.Gustos.Comidas)
	fmt.Printf("\n")

	
	//3. tambien podemos definir el campo de una persona de la siguiente manera
	var pers2 Persona

	pers2.Edad = 24

	fmt.Println(pers2)
}
