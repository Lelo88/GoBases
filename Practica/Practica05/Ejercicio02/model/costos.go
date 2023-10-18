package model

import "errors"

var errProductoInvalido = errors.New("producto invalido")

type Coste interface{
	GetCosto() float64
}

func NuevoCosto(producto string, precio float64) (Coste, error) {
	switch producto {
    case ProductoPequeño: 
	    return &Producto{
			Costo: precio,
		}, nil
	case ProductoMediano:
		return &Producto{
            Costo: precio,
			CostoMantenimiento: precio*0.03,
        }, nil
	case ProductoGrande:
		return &Producto{
            Costo: precio,
            CostoMantenimiento: precio*0.06,
			Shipping: 2500,
        }, nil 
    
	}
	return nil, errProductoInvalido
}
