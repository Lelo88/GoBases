package main

// no estaba el import de fmt
import "fmt"
var version = "3.1" // la variable global está mal definida de la siguiente manera version := "3.1"

func main() {
	name := "Golang" // Este string estaba definido así name := 'Golang', cuando debería ser name := "Golang"
	// ya que 'Golang' es un rune literal, no un string
	fmt.Println(name)
}
