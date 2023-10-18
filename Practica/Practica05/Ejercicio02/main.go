package main

import (
	"fmt"
	

	"github.com/Lelo88/GoBases/Practica/Practica05/Ejercicio02/model"
)

func main() {
	
	producto, err := model.NuevoCosto(model.ProductoGrande,100)
	
	if err!= nil {
        panic(err)
    }

	fmt.Println(producto)
	fmt.Println(producto.GetCosto())
}
