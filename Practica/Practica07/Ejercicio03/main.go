package main

import (
	"fmt"

	"github.com/Lelo88/GoBases/Practica/Practica07/Ejercicio03/clientes"
	"github.com/Lelo88/GoBases/Practica/Practica07/Ejercicio03/model"
)

func main() {
	defer func(){
		fmt.Println("Se finaliza la ejecucion del programa")
		if err := recover(); err!= nil {
			fmt.Println("Hay errores en la ejecucion de programa")
			fmt.Println(err)
		}
	}()

	cliente := model.Cliente {
		ID: 12345,
		Nombre: "Leandro Villalba",
		DNI: "M234234",
		Direccion: "Calle Falsa 123",
		Telefono: "12345123425",
	}

	guarda, err := clientes.NuevoAlmacenamiento()
	if err!= nil {
        panic(err)
    }

	cliente.ID = 12345
	err = guarda.Agregar(cliente)
	if err!= nil {
        panic(err)
    }

}

