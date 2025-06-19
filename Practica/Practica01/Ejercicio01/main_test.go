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
	var (
		nombre    = "Leandro"
		direccion = "Calle Falsa 123"
	)

	// Verifica que el valor de nombre sea "Leandro"
	assert.Equal(t, "Leandro", nombre, "El nombre debería ser Leandro")
	// Verifica que el valor de direccion sea "Calle Falsa 123"
	assert.Equal(t, "Calle Falsa 123", direccion, "La dirección debería ser Calle Falsa 123")
}

