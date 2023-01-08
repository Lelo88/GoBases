package main

import "fmt"

type Author struct {
	FirstName, LastName string
}

func (author Author) GetFullName() string{ //se apunta a la estructura autor para modificar el valor
	return fmt.Sprintf("%s %s", author.FirstName, author.LastName)
}

type Book struct{ 
	Title, Content string
	Author Author
}

//Hay dos maneras de realzar una interface

func (book Book) String() string{
	return fmt.Sprintf("El libro %s\ncontiene %s\n escrito por %s tiene un largo de %d\n", book.Title, book.Content, book.Author.GetFullName(), len(book.Content))
}

func main() {
	author := Author{
		FirstName: "John",
		LastName: "Doe",
	}

	book := Book{
		Title: "Como programar en Go",
		Content: "Lorem Ipsum...",
		Author: author, //le asigno el autor definido arriba
	}

	//en vez de tener un metodo que imprima todo, voy a crear una interface
	fmt.Println(book.String())

}

