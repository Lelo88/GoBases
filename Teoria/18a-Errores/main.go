package main

import "fmt"

func main() {
	error1 := fmt.Errorf("error1")
	error2 := fmt.Errorf("segundo error: %w", error1)

	fmt.Println(error2)
}

