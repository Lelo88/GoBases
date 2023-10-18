package main

import "fmt"

func main() {
	channel := make(chan string, 3)

	channel <- "Hola"
	channel <- "Mundo"
	channel <- "Teoria"

	fmt.Println(len(channel), cap(channel))

	close(channel) // cerrar el canal. ayuda a no crear exceso de datos en una goroutine

	for Conteo := range channel {
		fmt.Println(Conteo)
	}
}	
