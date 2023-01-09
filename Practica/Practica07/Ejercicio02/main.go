/*A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
Ahora que el archivo sí existe, el panic no debe ser lanzado.
Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.*/

package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err:= recover(); err!=nil{
			fmt.Println("Tipo de error: ", err)
		}
		fmt.Println("Finalizando ejecución...") //luego ejecuta esta accion al final 
	}()

	customers,err:= os.ReadFile("customers.txt")
	if err != nil {
		panic("No se encuentra el archivo")
	}

	fmt.Println(string(customers)) //-> primero ejecuta esta accion
}
