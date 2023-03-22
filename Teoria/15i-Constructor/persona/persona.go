package persona

type Persona struct {
	nombre string
	edad int
}

func (p *Persona) Constructor(name string, edad int){
	p.nombre = name
	p.edad = edad
}