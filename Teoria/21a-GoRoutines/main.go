//el paralelismo es mas rapido, pero es mas costoso. Por eso se enrutan tareas para que se aprovechen la concurrencia y
//el paralelismo.

//la palabra reservada go ejecuta en simultaneo un proceso

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func procesar(i int){
	rand := time.Duration(math.Round(rand.Float64()* 10)) //variable incializada para obtener luego un tiempo de proceso random
	fmt.Println(i, "-Inicia")
	time.Sleep(rand * time.Second)
	fmt.Println(i, "-Termina")
}

func main() {
	for i:=0;i<10;i++{
		go procesar(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("Termino el programa")
	//al quitarle las dos acciones de arriba, el programa va a finalizar antes de que las go routines se ejecuten. 
	//para eso vamos a agregarle channels
}
