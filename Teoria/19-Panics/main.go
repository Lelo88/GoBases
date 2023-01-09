//los panics son controladores de errores que cortan el programa de forma abrupta. Se utilizan para errores que NO deberian
//suceder para nada
//nos muestran en detalle donde ocurrio el error
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Iniciando...")
	_,err := os.Open("no-file.txt")
	if err!= nil {
        //panic(err)
		fmt.Println(err) //lo que hacemos ahora es imprimir el error y dar una salida con la funcion os.exit
		os.Exit(0) //lo mismo pasaria con return
    }

fmt.Println("Fin")	
}

