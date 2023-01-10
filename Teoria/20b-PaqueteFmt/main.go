package main

import "fmt"

type Estructura struct {
	Valor1 int
	Valor2 int
    Valor3 int
	Valor4 bool
	Estructura2
}

type Estructura2 struct {
	Campo1 int
	Campo2 int
}

func main() {

	es:=Estructura{}
	fmt.Printf("%v\n",es)
	fmt.Printf("%+v\n",es)
}

