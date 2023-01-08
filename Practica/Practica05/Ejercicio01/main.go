package main

import "fmt"

type Alumnos struct {
	Nombre, Apellido, DNI, Fecha string
}

type Detalle interface {
	Detalle() string
}

func GetAlumno(d Detalle){
	fmt.Println(d.Detalle())
}

func (alumno *Alumnos) Detalle() string{
	return alumno.Nombre + " " + alumno.Apellido + " " +  alumno.DNI + " " + alumno.Fecha
}

func main() {
	var alumnos Alumnos
    alumnos.Nombre = "Leandro"
    alumnos.Apellido = "Villalba"
    alumnos.DNI = "33902312"
	alumnos.Fecha = "30/07/1988"

	GetAlumno(&alumnos)
}


