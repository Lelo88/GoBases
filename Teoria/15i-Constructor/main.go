package main

import (
	"fmt"

	"github.com/Lelo88/GoBases/Teoria/15i-Constructor/persona"
)

func main() {
	p1 := persona.Persona{}
	p1.Constructor("Leandro", 34)
	fmt.Println(p1)
}