/*Un metodo es una funcion con un argumento receptor especial, definido con func y el nombre del metodo*/

package main

import (
	"fmt"
	"math"
)

//declaramos un struct circulo

type Circulo struct {
	radio float64
}

//aca crearemos el metodo para calcular el area de un circulo

func (c Circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

func main() {
	c := Circulo{10}
	c2:= Circulo{20}
	fmt.Printf("%.2f\t",c.area()) //al llamar al metodo, este realiza la accion que encomiendo hacerle, ya que tiene definido el campo radio
	fmt.Printf("%.2f\n",c.perim())


	fmt.Printf("%.2f\t",c2.area()) 
	fmt.Printf("%.2f\n",c2.perim())
}