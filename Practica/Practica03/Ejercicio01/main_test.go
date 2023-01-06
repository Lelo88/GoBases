package main 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculaImpuesto(t *testing.T){
	
	//inicializo las variables 
	sueldo1 := 40000.00
	sueldo2 := 55000.00
	sueldo3 := 162000.00

	//calculo los resultados esperados
	resultadoEsperado1 := 0.0
	resultadoEsperado2 := sueldo2 * 0.17
	resultadoEsperado3 := sueldo3 * 0.27

	//paso las variables para que la funcion me retorne el resultado 
	resultado1, ok := calculaImpuesto(sueldo1)
	resultado2, ok2 := calculaImpuesto(sueldo2)
	resultado3, ok3 := calculaImpuesto(sueldo3)

	//corroboro si la funcion responde como espero comparando los resultados
	assert.Equal(t, resultadoEsperado1, resultado1, ok)
	assert.Equal(t, resultadoEsperado2, resultado2, ok2)
	assert.Equal(t, resultadoEsperado3, resultado3, ok3)
}	