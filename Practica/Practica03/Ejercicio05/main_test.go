package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAnimal(t *testing.T) {
	perro, msg := animal("perro")
	cantidad := 10.0

	resultadoEsperado := 100.0
	
	resultado := perro(cantidad)

	assert.Equal(t, resultadoEsperado, resultado, msg)

}