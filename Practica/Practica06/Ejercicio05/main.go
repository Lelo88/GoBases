package main

import (
	"fmt"
	"errors"
)

func calculaSalario(horas int, valorHora float64) (float64, error) {

    if horas < 80 {
		return 0, errors.New("error: El empleado no puede haber trabajado menos de 80 horas mensuales")
	}

	if valorHora <= 0 {
		return 0, errors.New("el valor de la hora no es valido")
	}

	if horas*int(valorHora) > 150000 {
		return valorHora * float64(int(valorHora)) - valorHora * float64(int(valorHora)) * 0.1, nil
	}

	return valorHora * float64(horas), nil
}

func main() {
	salario, err := calculaSalario(80, 3.5)
	if err!= nil {
        panic(err)
    }

	fmt.Println(salario)

	/*salario2, err2 := calculaSalario(80, -3.5)

	if err2!= nil {
        panic(err2)
    }

	fmt.Println(salario2)*/

	salario3, err3 := calculaSalario(75, 3.5)

	if err3!= nil {
        panic(err3)
    }

	fmt.Println(salario3)
}