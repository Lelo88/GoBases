//los canales son mecanismos de comunicacion entre go routines. Son los mas simples de encontrar

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func procesar(i int, c chan int){
	rand := time.Duration(math.Round(rand.Float64()* 10))
	fmt.Println(i, "-Inicia")
	time.Sleep(rand * time.Second)
	fmt.Println(i, "-Termina")
	c <- i //le enviamos el valor de i al canal cuando se termina el proceso
}

func main() {
	
	c := make(chan int) //los channels siempre se definen asi
	
	for i:=0;i<10;i++{
		go procesar(i, c)
	}
	//time.Sleep(5000 * time.Millisecond)
	//fmt.Println("Termino el programa")
	//al quitarle las dos acciones de arriba, el programa va a finalizar antes de que las go routines se ejecuten. 
	//para eso vamos a agregarle channels

	for i:=0;i<10;i++{
		fmt.Println("Termino el programa en ", <-c)
	}
}