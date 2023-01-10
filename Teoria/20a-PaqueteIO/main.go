package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	s:= "Este texto es copiado\n"
	copia := strings.NewReader(s) //nos devuelve un puntero a una instancia de reader
	_,err := io.Copy(os.Stdout, copia) //lo envia a la terminal

	if err!= nil {
        log.Fatal(err)
    }

	cadena, err2:=io.ReadAll(copia)
	if err2!=nil {
        log.Fatal(err2)
    }
	fmt.Println(string(cadena))

	_,err3 := io.WriteString(os.Stdout, s)
	if err3 !=nil { 
		log.Fatal(err3)
	}	


}

