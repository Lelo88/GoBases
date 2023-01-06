package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculaPromedio(t *testing.T){
	notas := []int{4,5,6,3,-1}
	//ok1, notas2 := ingreso()
	//ninguno de los dos resultados deberia dar ok ya que tengo un numero mayor a 10 y uno menor a 0

	sumNotas1 := 0
	for _,value := range notas{
		sumNotas1+=value
	}

	/*sumNotas2 := 0
	for _,value := range notas2{
		sumNotas2+=value
	}*/

	resultadoEsperado1 := float64(sumNotas1)/float64(len(notas))
	//resultadoEsperado2 := float64(sumNotas2)/float64(len(notas2))

	resultado1 := calculaPromedio(notas...)
	//resultado2 := calculaPromedio(notas2...)

	assert.Equal(t, resultadoEsperado1, resultado1)
	//assert.Equal(t, resultadoEsperado2, resultado2, ok1)
}