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

func (c Circulo) GetArea() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) GetPerim() float64 {
	return 2 * math.Pi * c.radio
}

func (c Circulo) GetName() string {
	return "Circulo"
}
//agregado de clases interfaces
func printInformation(geometry Geometry) {
	fmt.Println("Geometry: ", geometry.GetName())
	fmt.Println("Circle area: ", geometry.GetArea())
	fmt.Println("Circle area: ", geometry.GetPerim())
}

//si queremos mas figuras y queremos usar mas detalles, vamos a implementar una interfaz general

type Geometry interface { //esta interfaz general implementa todas las funciones para las figuras que vaya creando
	GetName() string
	GetArea() float64
	GetPerim() float64
}

//A modo de ejemplo, creo otra figura

type Rectangle struct {
	Width  float64
	Hight  float64
}

func (r Rectangle) GetName() string {
    return "Rectangle"
}

func (r Rectangle) GetArea() float64 { 
	return r.Width * r.Hight
}

func (r Rectangle) GetPerim() float64 {
    return 2*(r.Hight + r.Width)
}	

func main() {
	c := Circulo{10}
	c2:= Circulo{20}
	
	printInformation(c)
    printInformation(c2)

	r1 := Rectangle{3,2}
	printInformation(r1)
}