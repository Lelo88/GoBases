package persona

type Persona struct {
	nombre string
	edad int
}

func (p *Persona) Constructor(name string, edad int){
	p.nombre = name
	p.edad = edad
}

func (p *Persona) Get() string{
	return p.nombre
}

func (p *Persona) Set(nombre string){
	p.nombre = nombre
}



