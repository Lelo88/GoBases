package main

import (
	"fmt"

	"github.com/Lelo88/GoBases/Teoria/15j-SetGet/persona"
)

func main() {
	per1 := persona.Persona{}
	
	per1.Constructor("Leandro", 34)
	
	fmt.Println("La persona se llama: " + per1.Get())
	per1.Set("Gabriel")
	fmt.Println("Ahora el nombre de la persona es: " + per1.Get())
}