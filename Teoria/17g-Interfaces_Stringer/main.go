package main

import "fmt"

type DistanceUnit int

const (
	Kilometer DistanceUnit = 0
	Mile      DistanceUnit = 1
)

type Distance struct {
	number float64
	unit   DistanceUnit
}

func (sc DistanceUnit) String() string {
	units := []string{"km", "mi"}
	return units[sc]
}

func (d Distance) String() string {
	return fmt.Sprintf("%v %v", d.number, d.unit)
}

func main() {
	kmUnit := Kilometer
	kmUnit.String()
	// => km

	mileUnit := Mile
	mileUnit.String()
	// => mi

	dist := Distance{
		number: 790.7,
		unit:   Kilometer,
	}

	// => 790.7 km

	fmt.Println(dist.String())
}
