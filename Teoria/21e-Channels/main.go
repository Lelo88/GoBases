package main

import "fmt"

func mensaje(msg string, channel chan string) {
	channel <- msg
}

func main() {
	channel := make(chan string)
	channel2 := make(chan string)

	go mensaje("Estoy en el canal 1", channel)
	go mensaje("Estoy en el canal 2", channel2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-channel:
			fmt.Println("Se ha ejecutado el canal 1", msg)
		case msg2 := <-channel2:
			fmt.Println("Se ha ejecutado el canal 2", msg2)
		}
	}
}