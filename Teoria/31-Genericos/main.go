package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// uso de any
// tipos y restricciones
// Crear restricciones
// Restricciones y operadores
// Estructuras

// uso de any

// any es un tipo de dato que puede ser cualquier tipo de dato
// se puede usar para hacer funciones que reciban cualquier tipo de dato
// o para hacer listas que puedan contener cualquier tipo de dato
func Printlist(list []any) {
	for _, v := range list {
		fmt.Println(v)
	}
}

// funciones genéricas
// se puede hacer una función que reciba cualquier tipo de dato
// se puede hacer una función que retorne cualquier tipo de dato
// con la virguilla ~ se puede determinar los tipos de datos que puede recibir la función de manera más restrictiva
func Sum[T Numbers](nums ...T) T {
	var result T
	for _, num := range nums {
		result += num
	}
	return result
}

// se pueden crear variables de otro tipo de dato y especificarlas en la función Sum, aunque no es correcto hacerlo

type integer int

// Así podemos construir restricciones para los tipos de datos que se pueden pasar a la función Sum
type Numbers interface{
	~int | ~float64 | ~float32 | ~uint
}

// Tambien se puede utilizar el paquete constraints para hacer restricciones. 

// Includes es una función que recibe una lista de cualquier tipo de dato y un valor de cualquier tipo de dato
// y retorna un valor booleano
func Includes[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}


func Filter[T constraints.Ordered](list []T, f func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, v := range list {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

type Product[T uint | string] struct {
	ID T
	Name string
	Price float64
}

func main() {
	var num1 integer = 100
	var num2 integer = 200
	Printlist([]any{1, 2, 3, 4, 5})
	Printlist([]any{"hola", "mundo", "como", "estas"})
	fmt.Printf("\n")
	fmt.Println("La suma total es: ",Sum(1, 2, 3, 4, 5, 6.6))
	fmt.Println(Sum(num1, num2, 2, 3, 4))
	fmt.Println(Includes([]int{1, 2, 3, 4, 5}, 3))
	strings := []string{"hola", "mundo", "como", "estas"}
	fmt.Println(Includes([]string{"hola", "mundo", "como", "estas"}, "hola"))
	fmt.Println(Includes([]string{"hola", "mundo", "como", "estas"}, "adios")) // no está en la lista, es false
	fmt.Println(Includes([]float64{1.1, 2.2, 3.3, 4.4, 5.5}, 3.3))
	fmt.Println(Includes(strings, "hola"))
	fmt.Printf("\n")
	// Filter
	// se puede hacer una función que filtre una lista de cualquier tipo de dato
	// se le pasa una lista y una función que recibe un valor de la lista y retorna un valor booleano
	// la función retorna una lista con los valores que cumplan con la condición de la función
	fmt.Println(Filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 2
	}))

	fmt.Println(Filter(strings, func(s string) bool {
		return len(s) > 4
	}))

	fmt.Printf("\n")
	// Estructuras
	// se pueden hacer estructuras genéricas
	// se pueden hacer restricciones para los tipos de datos que se pueden pasar a la estructura
	// se pueden hacer estructuras que reciban cualquier tipo de dato
	// se pueden hacer estructuras que retornen cualquier tipo de dato
	// se pueden hacer estructuras que reciban cualquier tipo de dato y retornen cualquier tipo de dato
	// se pueden hacer estructuras que reciban cualquier tipo de dato y retornen cualquier tipo de dato con restricciones
	// se pueden hacer estructuras que reciban cualquier tipo de dato y retornen cualquier tipo de dato con restricciones y operadores

	producto1 := Product[uint]{ID: 1, Name: "Producto 1", Price: 100.0}
	producto2 := Product[string]{ID: "2", Name: "Producto 2", Price: 200.0}
	fmt.Println(producto1)
	fmt.Println(producto2)
}