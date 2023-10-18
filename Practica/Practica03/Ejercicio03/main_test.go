package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculaSalario(t *testing.T) {
		categoria := "A"
		horas := 45

		resultadoEsperado := float64(horas * 3000)
		resultadoEsperado += resultadoEsperado * 0.5

		resultado := calculaSalario(categoria, horas)

		assert.Equal(t, resultadoEsperado, resultado)
}