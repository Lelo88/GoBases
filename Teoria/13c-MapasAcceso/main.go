package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Print(a["brand"]) // Ford
	a["year"] = "2007" // Actualizar el valor de una clave
	fmt.Print(a["year"]) // 2007
}
