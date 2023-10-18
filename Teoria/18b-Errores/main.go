//paquete errores: contiene 4 funciones: new, unwrap, is, as

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	err:= errors.New("este es un error")
	fmt.Println(err)

	err1 := fmt.Errorf("primer error")
	err2 := fmt.Errorf("segundo error: %w", err1)

	err3 := errors.Unwrap(err2)
	if err3 == err1 {
		fmt.Println("Mismo error")
	}

	fmt.Printf("\n")

	_,errorIs := os.Open("not_exists.txt")
	var notExist error = fs.ErrNotExist
	if errors.Is(errorIs, notExist){ //is compara dos errores y verifica si son iguales
		fmt.Println("file not exist")
	}

	fmt.Printf("\n")

	_,errorAs := os.Open("not_exists.txt")
	var pathError *fs.PathError
	if errors.As(errorAs, &pathError) { //as compara dos errores y devuelve un booleano
		fmt.Println("Path Error", pathError.Path)
	}

}