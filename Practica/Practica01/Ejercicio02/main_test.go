package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"
// TestMain es la función principal de prueba que se ejecuta antes de cualquier otra prueba
func TestMain(m *testing.M) {
	fmt.Println("Iniciando las pruebas...")
	// Aquí podrías realizar configuraciones iniciales si es necesario
	m.Run() // Ejecuta las pruebas
	fmt.Println("Pruebas finalizadas.")
}
// TestVariables es una prueba unitaria que verifica el valor de las variables nombre y direccion
func TestVariables(t *testing.T) {
	t.Run("TestVariables", func(t *testing.T) {
		var temperatura float64
		var humedad uint8
		var presion uint16

		temperatura = 30.6
		humedad = 75
		presion = 1084

		expectedOutput := "La temperatura en el dia de hoy es de 30.6, con una humedad de 75 y una presion de 1084"
		actualOutput := fmt.Sprintf("La temperatura en el dia de hoy es de %.1f, con una humedad de %d y una presion de %d", temperatura, humedad, presion)

		assert.Equal(t, expectedOutput, actualOutput, "Los valores no coinciden")
	})

	t.Run("TestTemperatura", func(t *testing.T) {
		var temperatura float64 = 30.6
		assert.Greater(t, temperatura, 0.0, "La temperatura debe ser mayor que 0")
	})
	
	t.Run("TestHumedad", func(t *testing.T) {
		var humedad uint8 = 75
		assert.GreaterOrEqual(t, humedad, uint8(0), "La humedad debe ser mayor o igual a 0")
	})

	t.Run("TestPresion", func(t *testing.T) {
		var presion uint16 = 1084
		assert.Greater(t, presion, uint16(0), "La presión debe ser mayor que 0")
	})
	}