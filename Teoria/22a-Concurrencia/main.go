//las concurrencias son procesos que estan alineados con multiples procesos en un determinado tiempo
// paralelismo es un proceso que esta haciendo multiples cosas al mismo tiempo

package main

import (
	"fmt"
	"time"
)

func MuchoTexto (Palabra string) {
	fmt.Println(Palabra)
}

func main() {
	MuchoTexto("Hola a todos")
	// al agregar el prefijo go, el programa se ejecuta y realiza una concurrencia
	// pero se ejecuta a la vez

	go MuchoTexto("Esta es una goroutine")
	
	time.Sleep(time.Second * 2)
} 