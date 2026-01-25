package main

import "fmt"

// declaring a new struct type
type car struct {
	brand string
	price int
}

// declaring a method for a car type
// this method receives a car by value
func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// declaring another method for a car type
// this method receives a car by pointer
func (c *car) changeCar2(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

type distance *int

func main() {
	myCar := car{
		brand: "Ford",
		price: 20000,
	}
	
	fmt.Println(myCar)
	
	myCar.changeCar1("Toyota", 30000)
	fmt.Println(myCar)
	
	myCar.changeCar2("BMW", 40000)
	fmt.Println(myCar)
}
