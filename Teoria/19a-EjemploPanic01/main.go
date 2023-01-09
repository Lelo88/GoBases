package main

import "fmt"

func ejemplo() {
	animales :=[] string {
		"vaca",
		"perro",
		"halcon",
		"oso",
	}

	fmt.Println(animales[5]) // vamos a generar un error a proposito para que en la terminal me tire un panic
}

func main() {
	fmt.Println("INICIANDO")


	ejemplo()

    fmt.Println("FINALIZANDO")
}

