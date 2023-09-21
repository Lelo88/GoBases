package main

import "fmt"

type error interface {
	Error() string
}

// errorstring es una implementacion trivial de errores
type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

// se puede construir un error
func New(text string) error {
	return errorString{text}
}

func main() {
	// New("error")
	newError := New("este es un error creado por mi")
	fmt.Println(newError)
}
