package clientes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

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
	info, err := os.ReadFile(archivo.NombreArchivo)
	if err!= nil {
        return fmt.Errorf("el archivo %s no existe", err)
    }

	err =json.Unmarshal(info, &archivo)
	if err!= nil { 
		return fmt.Errorf("el archivo %s no existe", err)
	}
	
	return nil
}

func (archivo *GuardaArchivo) EscribirArchivo() error { 
	
	info, err := json.Marshal(archivo)
	if err!= nil { 
		return fmt.Errorf("el archivo %s no existe", err)
    }

	err =os.WriteFile(archivo.NombreArchivo, info,os.FileMode(0644))
	
	if err!= nil { 
		return fmt.Errorf("el archivo %s no existe", err)
    }

	return nil
}

func (archivo *GuardaArchivo) Agregar(cliente model.Cliente) error {
	err := cliente.ValidaCliente()
	if err!= nil {
        return err
    }

	err = archivo.ValidaClienteArchivo(cliente)

	if err!= nil {
        return err
    }

	archivo.Clientes = append(archivo.Clientes, cliente)

	err = archivo.EscribirArchivo()
	if err!= nil {
        return err
    }

    return nil
}

func NuevoAlmacenamiento() (GuardarCliente, error){
	guarda := &GuardaArchivo{
		NombreArchivo: "./clientes.json",
	}

	err:= guarda.LeerArchivo()
	if err!=nil{
        return nil,err
    }

	return guarda, nil
}

