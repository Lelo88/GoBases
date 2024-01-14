package main

import (
	"fmt"
	"math"
)

const (
	RectangleType = "RECTANGLE"
	CircleType    = "CIRCLE"
)

type Geometry interface {
	GetName() string
	GetArea() float64
	GetPerimeter() float64
}

type Rectangle struct {
	Width float64
	High  float64
}

// GetPerimeter implements Geometry
func (r Rectangle) GetPerimeter() float64 {
    return 2 * (r.High + r.Width)

}

func (r Rectangle) GetName() string {
	return "Rectangle"
}

func (r Rectangle) GetArea() float64 {
	return r.Width * r.High
}



type Circle struct {
	Radius float64
}

// GetArea implements Geometry
func (c Circle) GetArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

// GetName implements Geometry
func (c Circle) GetName() string {
	return "Circle"
}

// GetPerimeter implements Geometry
func (c Circle) GetPerimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//constructor 
func NewGeometry(name string, value ...float64) (instance Geometry, ok bool) {
	switch name {
	case RectangleType:
		instance = Rectangle{
			Width: value[0],
			High:  value[1],
		}
		ok = true
	case CircleType:
		instance = Circle{
			Radius: value[0],
		}
		ok = true

	}
	return
}

func medidas (g Geometry){
	fmt.Println("Nombre de la figura: " + g.GetName())
	fmt.Printf("Perimetro: %.2f\n",g.GetPerimeter())
	fmt.Printf("Area: %.2f\n",g.GetArea())
}

func main() {
	geometry, ok := NewGeometry(RectangleType,  2, 3)
    if !ok {
        panic("failed to create geometry")
    }
    medidas(geometry)
}