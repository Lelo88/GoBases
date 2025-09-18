package main

import "testing"

func Test_AdivinaNumero(t *testing.T) {
	tests := []struct {
		numero    int
		mensaje   string
		expectErr bool
	}{
		{7, "¡Felicidades! Has adivinado el número.", false},
		{5, "Lo siento, no has adivinado el número.", false},
		{0, "", true},
		{11, "", true},
	}

	for _, test := range tests {
		resultado, err := AdivinaNumero(test.numero)
		
		if (err != nil) != test.expectErr {
			t.Errorf("Error inesperado para el numero %d: %v", test.numero, err)
			continue
		}

		if resultado != test.mensaje {
			t.Errorf("Para el numero %d, se esperaba '%s', pero se obtuvo '%s'", test.numero, test.mensaje, resultado)
		}
	}
}

