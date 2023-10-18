package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestOperacion(t *testing.T) {
	notas := []float64{4,5,6,3,2,1}
	min := operacionProfesores(minimum)

	resultadoEsperado := 1.0

	resultado := min(notas...)

	assert.Equal(t, resultadoEsperado, resultado)
}