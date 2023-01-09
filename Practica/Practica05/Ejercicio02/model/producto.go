package model

const (
	ProductoPequeño = "pequeño"
	ProductoMediano = "mediano"
	ProductoGrande   = "grande"
)

type Producto struct {
	Costo, CostoMantenimiento, Shipping float64
}

func (p *Producto) GetCosto() float64 {
	return p.Costo + p.CostoMantenimiento + p.Shipping
}