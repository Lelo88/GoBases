package model

import (
	"errors"
	"fmt"
)

var errValidaCliente = errors.New("el cliente tiene campos con valor 0")

type Cliente struct {
    ID 	   uint   		`json:"id"`
	Nombre string 		`json:"nombre"`
	DNI    string 		`json:"dni"`
	Telefono string 	`json:"telefono"`
	Direccion string 	`json:"direccion"`
}

func (cliente *Cliente) ValidaCliente() error {
	if ok:= cliente.ValidaDatos(); !ok {
		return fmt.Errorf("error en validación de cliente: %w", errValidaCliente)
	}
	return nil
}

func (cliente Cliente) ValidaDatos() bool {
	return cliente.DNI!= "" && cliente.Telefono!= "" && cliente.Direccion!= "" && cliente.ID != 0 && cliente.Nombre!= ""
}