package main

import "fmt"

func mensaje(msg string, channel chan string) {
	channel <- msg
}

func main() {
	channel := make(chan string, 1)

	fmt.Println("Hola a todos")

	go mensaje("Estoy en el canal", channel)

	fmt.Println(<-channel)
}