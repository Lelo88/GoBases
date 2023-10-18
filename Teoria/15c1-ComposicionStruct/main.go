package main

import "fmt"

type Author struct {
	FirstName, LastName string
}

func (author *Author) GetFullName() string{ //se apunta a la estructura autor para modificar el valor
	return fmt.Sprintf("%s %s", author.FirstName, author.LastName)
}

type Book struct{ 
	Title, Content string
	Author Author
}

func (book *Book) PrintInformation(){
	fmt.Printf("Libro: %s\n", book.Title)
	fmt.Printf("Cantidad de caracteres: %d\n", len(book.Content))
	fmt.Printf("Autor: %s\n", book.Author.GetFullName())
}

func main(){
	author := Author{
		FirstName: "John",
		LastName: "Doe",
	}

	book := Book{
		Title: "Como programar en Go",
		Content: "Lorem Ipsum...",
		Author: author, //le asigno el autor definido arriba
	}
	
	book.PrintInformation()
}