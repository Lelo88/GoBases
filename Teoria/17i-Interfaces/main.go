package main

import "fmt"

// Utilizamos interfaces para aislar y desacoplar una parte del código
// de otras partes del código
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * 3.14
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) volume() float64 {
	return c.radius * c.radius * c.radius * 4 / 3 * 3.14
}

func printShapeInfo(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}


func main() {
	c := circle{radius: 5}
	r := rectangle{width: 5, height: 10}
	
	printShapeInfo(c)
	printShapeInfo(r)

	// que tipo me marca interface

	var shape1 shape
	fmt.Printf("Type: %T\n", shape1) // nil
	shape1 = c
	fmt.Printf("Shape type: %T\n", shape1) // main.circle
	shape1 = r
	fmt.Printf("Shape type: %T\n", shape1) // main.rectangle

	var shape2 shape = circle{radius: 10}
	fmt.Printf("Shape2 type: %T\n", shape2) // main.circle

	// fmt.Println("Volume:", shape2.(circle).volume())

	ball, ok := shape2.(circle)
	if ok == true {
		fmt.Println("Volume:", ball.volume())
	}

	switch value := shape2.(type) {
	case circle:
		fmt.Printf("%#v has circle type.", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type.", value)
	}
}



