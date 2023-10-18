/*Como la herencia no existe en GO, existen las estructuras embebidas, que son similares. 
La composicion en Go se encarga de crear programas mas grandes a traves de partes mas pequeñas. */

package main

import "fmt"

type Vehiculo struct { //estructura padre
	Km 		float64
	Tiempo  float64
}

type Auto struct { //estructura hija (embebida)
	Vehiculo
}

type Moto struct { //estructura hija (embebida)
	Vehiculo
}

//metodo de vehiculo
func (v Vehiculo) DetalleVehiculo(){
	fmt.Printf("km: \t %.2f\ntiempo:\t%.2f\n", v.Km,v.Tiempo)
}

//metodos de auto 
func (a *Auto) Correr(minutos int){
	a.Tiempo = float64(minutos) / 60
	a.Km = a.Tiempo * 100
}

func (a *Auto) Detalle(){
	fmt.Println("\nV:\tAuto")
	a.DetalleVehiculo()
}

func (m *Moto) Correr(minutos int){
	m.Tiempo = float64(minutos) / 60
	m.Km = m.Tiempo * 100
}

func (m *Moto) Detalle(){
	fmt.Println("\nV:\tMoto")
	m.DetalleVehiculo()
}


func main(){
	a := Auto{}
	a.Correr(123)
	a.Detalle()

	m := Moto{} 
	m.Correr(100)
	m.Detalle()
}