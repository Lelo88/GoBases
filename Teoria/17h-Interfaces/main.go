package main

import (
	"github.com/Lelo88/GoBases/Teoria/17h-Interfaces/animal"
	"github.com/Lelo88/GoBases/Teoria/17h-Interfaces/book"
)

func main(){
	var myBook = book.NewBook("The Hobbit", "J.R.R. Tolkien", 293)
	var myTextBook = book.NewTextBook("The Go Programming Language", "Alan A. A. Donovan & Brian W. Kernighan", "Addison-Wesley", "Advanced", 380)
	
	myBook.PrintInfo()
	myTextBook.PrintInfo()

	myBook.SetTitle("The Lord of the Rings")
	myBook.PrintInfo()


	animales := []animal.Animal{
		&animal.Perro{Nombre: "Firulais"},
		&animal.Gato{Nombre: "Garfield"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}