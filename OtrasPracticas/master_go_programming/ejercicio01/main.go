package main

import "fmt"

func main() {
	var (
		a int
		b string
		c bool
		d float64
	)

	x, y, z := 20, 15.5, "Gophers"
	score := 99

	fmt.Printf("a=%d b=%s c=%t d=%f x=%d y=%f z=%s score=%d\n", a, b, c, d, x, y, z, score)
	fmt.Printf("z=%q\n", z)
	fmt.Printf("a=%v b=%v c=%v d=%v x=%v y=%v z=%v score=%v\n", a, b, c, d, x, y, z, score)
	fmt.Printf("y type=%T\n", y)
	fmt.Printf("score type=%T\n", score)
}
