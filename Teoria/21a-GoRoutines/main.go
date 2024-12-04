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

// La concurrencia en golang se logra con la palabra reservada go. Se van ejecutando procesos en simultaneo.
// En este caso, se ejecutan 10 procesos en simultaneo, pero el programa finaliza antes de que se ejecuten todos los procesos.
// Para eso se utilizan los channels.
// Un channel es un tipo de dato que permite la comunicación entre goroutines.
// Se pueden enviar y recibir datos a través de un channel.
// Se crean con la función make() y se utilizan con la flecha <-
// En este caso, se va a crear un channel de tipo int, y se va a enviar el valor i a través del channel.
// Luego, se va a recibir el valor del channel y se va a imprimir.
// Se va a utilizar un channel para que el programa espere a que todas las goroutines terminen de ejec
//utar antes de finalizar el programa.

