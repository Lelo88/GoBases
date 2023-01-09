package clientes

import (
	"errors"

	"github.com/Lelo88/GoBases/Practica/Practica07/Ejercicio03/model"
)

var errClienteExiste = errors.New("cliente existente")

type GuardarCliente interface {
	Agregar(model.Cliente) error
}

type GuardaArchivo struct {
	NombreArchivo   string
	Clientes []model.Cliente `json:"clientes"`	
}

func (archivo *GuardaArchivo) ValidaClienteArchivo(cliente model.Cliente) error {
	for _, cliente2 := range archivo.Clientes {
        if cliente.ID == cliente2.ID {
            return errClienteExiste
        }
    }
    return nil
}

func (archivo *GuardaArchivo) LeerArchivo() error {

	return nil
}

func (archivo *GuardaArchivo) Agregar(cliente model.Cliente) error {
	err := cliente.ValidaCliente()
	if err!= nil {
        return err
    }
	return nil	
}



