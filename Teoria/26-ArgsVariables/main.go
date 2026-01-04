package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const requiredArgs = 3
	const maxArgLen = 32
	const maxTotalLen = 96

	args := os.Args[1:]
	argNames := []string{"nombre1", "nombre2", "apellido"}
	var errs []string

	if len(args) < requiredArgs {
		errs = append(errs, fmt.Sprintf("faltan argumentos: se esperaban %d y se recibieron %d", requiredArgs, len(args)))
	}
	if len(args) > requiredArgs {
		errs = append(errs, fmt.Sprintf("sobran argumentos: se esperaban %d y se recibieron %d (extra: %q)", requiredArgs, len(args), args[requiredArgs:]))
	}

	checkN := len(args)
	if checkN > requiredArgs {
		checkN = requiredArgs
	}

	totalLen := 0
	for i := 0; i < checkN; i++ {
		trimmed := strings.TrimSpace(args[i])
		args[i] = trimmed
		if trimmed == "" {
			errs = append(errs, fmt.Sprintf("%s no puede estar vacío", argNames[i]))
			continue
		}
		if len(trimmed) > maxArgLen {
			errs = append(errs, fmt.Sprintf("%s supera el tamaño máximo (%d): %d", argNames[i], maxArgLen, len(trimmed)))
		}
		totalLen += len(trimmed)
	}
	if totalLen > maxTotalLen {
		errs = append(errs, fmt.Sprintf("la suma de tamaños de los argumentos supera el máximo (%d): %d", maxTotalLen, totalLen))
	}

	if len(errs) > 0 {
		fmt.Println("Errores:")
		for _, err := range errs {
			fmt.Println("-", err)
		}
		fmt.Println("Uso: ./greeter <nombre1> <nombre2> <apellido>")
		fmt.Println("Recibidos:", len(os.Args)-1)
		return
	}

	fmt.Printf("%#v\n", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("First argument:", args[0])
	fmt.Println("Second argument:", args[1])
	fmt.Println("Third argument:", args[2])
	fmt.Println("Number of arguments:", len(os.Args))
	fmt.Println("Hola")
}

/*
 % go build -o greeter (se debe ejecutar cada vez que se modifique el codigo)
% ./greeter hi leandro villalba
[]string{"./greeter", "hi", "leandro", "villalba"}Path: ./greeter
First argument: hi
Second argument: leandro
Third argument: villalba
Number of arguments: 4
*/
