package main

import (
	"fmt"
	"os"
	"log"
)
func main() {
	leerArchivo()
	crearArchivo()
	caracteristicas()
}

func leerArchivo(){
	data, err := os.ReadFile("./bootcampGo.txt")
	if err!= nil {
        log.Fatal(err)
    }
	fmt.Println(data)
	fmt.Println(string(data)) //si no le agrego el string, me va a imprimir los caracteres byte por byte
}

func crearArchivo(){ //creamos un archivo desde nuestro programa 
	archivo, err := os.Create("./bootcampGo2.txt")
	if err!= nil {
        log.Fatal(err)
    }

	_, err2 := archivo.WriteString(string("Este texto proviene de nuestro programa"))
	if err2!= nil {
        log.Fatal(err2)
    }

	fmt.Println("Hecho")
}

func caracteristicas(){ 
	archivo:="./bootcampGo.txt"
	data, err := os.Stat(archivo)

	if err!= nil {
        log.Fatal(err)
    }
	
	fmt.Println("Es un directorio: ", data.IsDir())
	fmt.Println("Nombre del archivo: ",data.Name())
	fmt.Println("Tamaño del archivo: ",data.Size())
	fmt.Println("Fecha y hora del archivo: ",data.ModTime())
	fmt.Println("Permisos del archivo: ",data.Mode())
}
