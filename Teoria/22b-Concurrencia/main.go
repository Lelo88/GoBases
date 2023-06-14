// waitgroup

package main

import (
	"fmt"
	"sync"
	"time"
)

func MuchoTexto (Palabra string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(Palabra)
}

func main() {
	fmt.Println("Hola a todos")

	var wg sync.WaitGroup
	
	wg.Add(1)

	go MuchoTexto("Esta es otra goroutine", &wg)
	wg.Wait()

	go func() {
		fmt.Println("Hasta la próxima")		
	}()

	time.Sleep(time.Second * 1)
} 