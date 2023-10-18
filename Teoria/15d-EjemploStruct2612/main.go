package main

import "fmt"

type Persona struct{ //estructur padre
	FirstName 	string
	LastName 	string
	Age 		uint
	Location 	
}

type Location struct{ //estructura de estructura padre
	Country		string
	Latitud		float64
	Longitud	float64
}

func (l *Location) SetLatitud(lat float64){ //si no agrego el asterisco, no va a cambiar el campo latitud que inicialice al principio
	l.Latitud = lat
}

func (l *Location) SetLongitud(lon float64){ //idem metodo seteo de latitud
	l.Longitud = lon
}

func (p Persona) FullName() string {
	return p.FirstName + " " + p.LastName
}

type Medico struct { //estructura embebida de padre
	Persona
	Matricula int
}

func (m *Medico) FalseSetMedico(matricula int){
	m.Matricula = matricula
}

type Cocinero struct{ //estructura embebida de padre 
	Persona 
}

func main(){
	medico := Medico{
		Persona{
			FirstName: "Alejandro",
			LastName: "Rodriguez",
			Age: 45, 
			Location: Location{
				Country: "Argentina",
				Latitud: 123.123,
				Longitud: 12312.123,
			},
		}, 21342,//matricula
	}
	//hemos creado un objeto medico a partir de la estructura medico, que esta conformada 
	fmt.Println(medico.Age)

	medico.SetLatitud(435.323)
	fmt.Println(medico.Latitud)
	fmt.Println(medico.FullName())
	fmt.Println(medico.Matricula)
	fmt.Println()

}