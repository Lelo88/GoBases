//el siguiente programa simula una funcionalidad que accede a una bbdd o envia una informacion a una API
package main

import (
	"fmt"
	"time"
)

func procesar(i int){
	fmt.Println(i, "-Inicia")
	time.Sleep(2 * time.Second)
	fmt.Println(i, "-Termina")
}

func main() {
	for i:=0;i<10;i++{
		procesar(i)
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Termino el programa")
}

