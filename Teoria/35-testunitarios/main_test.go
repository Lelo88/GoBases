package main

import (
	"testing"
)

func TestSumar(t *testing.T) {
	result := Sumar(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Sumar(2, 3) = %d; want %d", result, expected)
	}
}

func TestMayor(t *testing.T) {
	result := Mayor(10, 5)
	expected := 10
	if result != expected {
		t.Errorf("Mayor(10, 5) = %d; want %d", result, expected)
	}

	result = Mayor(3, 7)
	expected = 7
	if result != expected {
		t.Errorf("Mayor(3, 7) = %d; want %d", result, expected)
	}
}

func TestSumar_TableCases(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-2, -3, -5},
	}

	for _, tt := range tests {
		result := Sumar(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Sumar(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestMayor_TableCases(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{10, 5, 10},
		{3, 7, 7},
		{0, 0, 0},
		{-1, -5, -1},
	}

	for _, tt := range tests {
		result := Mayor(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Mayor(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

// cambiar los valores esperados para ver cuanto tarda la ejecucion.
// Las instrucciones son: 

//  1 go test
// 2 go test -cover
// 3 go test -coverprofile=coverage.out
// 4 go tool cover -html=coverage.out
// go test -cpuprofile=cpu.out && go tool pprof cpu.out
// una vez adentro del perfilador, usar el comando "web" para ver el perfil en el navegador
// o "top" para ver las funciones que más tiempo consumen en la terminal
// o "list NombreFuncion" para ver el detalle de una función específica
// o pdf para generar un archivo pdf con el perfil
func TestFibonacci(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13}, // Caso más grande para probar eficiencia
	}

	for _, tt := range tests {
		result := Fibonacci(tt.n)
		if result != tt.expected {
			t.Errorf("Fibonacci(%d) = %d; want %d", tt.n, result, tt.expected)
		}
	}
}

