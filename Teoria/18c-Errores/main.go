//errores devueltos en funciones

package main

import (
	"errors"
	"fmt"
)

func sayHi(name string) (string, error) { //cuando realizamos esta funcion, verificamos que retornemos dos tipos de variables
	if name == "" {
		return "", errors.New("name is empty")
	}
	return fmt.Sprintf("Hi, %s", name), nil
}

type myError struct {
	msg, x string	
}

func (e *myError) Error() string {
    return fmt.Sprintf("ha ocurrido un error: %s, %s", e.msg, e.x)
}

func main() {
	grettings, err := sayHi("") //si envio un string vacio, me retorna el error, caso contrario, el saludo
	if err!= nil {
        fmt.Println(err)
    }

	fmt.Println(grettings)

	fmt.Println("---------------------------")

	e := &myError{"nuevo error", "404"}

	var e2 *myError

	IsMyError := errors.As(e, &e2) //compara los errores
	fmt.Println(IsMyError)
}