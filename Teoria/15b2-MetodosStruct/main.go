package main

import "fmt"

type Persona struct {
	nombre string
	edad int
}

func (p *Persona) imprimir(){
	fmt.Printf("Nombre: %s, edad: %d \n", p.nombre, p.edad)
}

func (p *Persona) editarNombre(nuevoNombre string){
	p.nombre = nuevoNombre
}

func main(){
	persona1 := Persona {"Leandro", 34}
	persona1.imprimir()
	persona1.editarNombre("Gabriel")
	persona1.imprimir()

	persona2 := Persona{"Romina", 36}
	persona2.imprimir()
	persona2.editarNombre("Mara")
	persona2.imprimir()


}