package main 

import "fmt"

//interface
type OperacionAritmetica interface{
	resultado() int64
}

//struct operacion suma
type suma struct{
	numero1 int64
	numero2 int64
}

func (s suma) resultado() int64{
	return s.numero1 + s.numero2
}

// struct operacion resta
type resta struct{
	numero1 int64
	numero2 int64
}

func (s resta) resultado() int64{
	return s.numero1 - s.numero2
}

func Operacion (o OperacionAritmetica) {
	fmt.Println("El resultado de la operación es: ", o.resultado())
}

func main() {
	Operacion(suma{numero1: 10, numero2: 20})
	Operacion(resta{numero1: 10, numero2: 20})
}	