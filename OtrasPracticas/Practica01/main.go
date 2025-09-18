// voy a realizar un programa de practica, donde eligiendo un numero del 1 al 10, se adivinara.
// Voy a implementar test driven development para hacerlo

package main 

import "errors"

var ErrorNumeroNoValido = errors.New("el numero debe estar entre 1 y 10")

func AdivinaNumero(numero int) (string, error){
	if numero < 1 || numero > 10 {
		return "", ErrorNumeroNoValido
	}

	if numero == 7 {
		return "¡Felicidades! Has adivinado el número.", nil
	}

	return "Lo siento, no has adivinado el número.", nil
} 

func main() {
	// Este es el punto de entrada del programa, pero no es necesario para los tests.
	// Aquí podrías implementar una interfaz de usuario o lógica adicional si lo deseas.
	// Por ahora, solo se ejecutarán los tests.
}