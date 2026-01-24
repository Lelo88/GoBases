package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := 99

	fmt.Printf("x=%d y=%f z=%s score=%d\n", x, y, z, score) // x=10 y=15.500000 z=Gophers score=99
	fmt.Printf("z=%q\n", z)                                // z="Gophers"
	fmt.Printf("x=%v y=%v z=%v score=%v\n", x, y, z, score) // x=10 y=15.5 z=Gophers score=99
	fmt.Printf("y type=%T\n", y)                           // y type=float64
	fmt.Printf("score type=%T\n", score)                   // score type=int
}
