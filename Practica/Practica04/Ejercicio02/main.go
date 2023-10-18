package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	ID 			int
	Name 		string
	DateOfBirth string
}

type Empleado struct {
	Person
	ID		 int
	Position string
}

func (e Empleado) PrintEmployee() {
	id := strconv.Itoa(e.ID)
	fmt.Println( id + " " + e.Name + " " + e.Position + " " + e.DateOfBirth)  
}

func main(){
	persona := Person{
		ID: 12312,
		Name: "Alejandro Rodriguez",
		DateOfBirth: "12/03/1977",
	}

	empleado := Empleado{
		Person: 
			Person{
				ID: 545454,
				Name: "Leandro Villalba",
				DateOfBirth: "30/07/1988",
			},
		ID: 213423,
		Position: "Software Developer",
	}

	empleado.PrintEmployee()

	fmt.Println(persona)
	
}