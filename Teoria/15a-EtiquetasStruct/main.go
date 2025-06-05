/*Al definir una estructura, puedo añadirle etiquetas para que se identifiquen con variables de archivos con otro formato*/
/*Por lo general se trabaja con Json, formato que estructura datos*/
package main

import (
	"encoding/json"
	"fmt"
)

/*type MiEstructura struct {
	Campo1 string `miEtiqueta:"valor"`
	Campo2 string `miEtiqueta:"valor"`
	Campo3 string `miEtiqueta:"valor"`
}*/

type Persona struct{
	Nombre 		string `json:"nombre"`
	Apellido 	string `json:"apellido"`
	Telefono 	string `json:"telefono"`
	Direccion 	string `json:"direccion,omitempty"`
}
//En el caso de que nos encontremos con un campo vacio en un archivo json, le agregamos el campo omitempty, como se detalla arriba
//En el caso de que querramos ignorar datos que no nos interese, directamente colocaremos un guion al lado del campo:  `json:"-"`
func main(){
	per := Persona {"Leandro","Villalba","123456","Calle Falsa 123"}

	miJson, err := json.Marshal(per)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}

	fmt.Println(string(miJson))
	fmt.Println(err)
}