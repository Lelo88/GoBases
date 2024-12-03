package book

import "fmt"

//  un interface es un contrato que garantiza que un tipo concreto va a tener un conjunto de métodos.
type Printable interface {
	PrintInfo()
}

func Print(p Printable){
	p.PrintInfo()
}

type Book struct {
	title string //campos privados
	author string
	pages int
}

// Constructor de Book. Recibe todos los campos para crear un nuevo objeto Book.
func NewBook(title, author string, pages int) *Book {
	return &Book{
		title: title,
		author: author,
		pages: pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo(){
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

// TextBook es un tipo que hereda de Book. Es una composición de Book.
type TextBook struct {
	Book
	editorial string
	level string
}

// Constructor de TextBook. Recibe todos los campos para crear un nuevo objeto TextBook.
func NewTextBook(title, author, editorial, level string, pages int) *TextBook {
	return &TextBook{
		Book: Book{
			title: title,
			author: author,
			pages: pages,
		},
		editorial: editorial,
		level: level,
	}
}

func (t *TextBook) PrintInfo(){
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n", t.title, t.author, t.pages, t.editorial, t.level)
}

