package main

import "fmt"

func main(){
	fmt.Println("La suma de 4 y 5 es de",suma(4,5)) //una manera de llamar a la funcion es de manera directa 

	resultado := suma(6,7) //otra manera de llamar a una funcion, asociandola a una variable
	fmt.Println("El resultado entre 6 y 7 es", resultado)

	resultado = sumaVariadica(1,2,3,4,5,6,7,8,9,10) //llamamos a la funcion variadica
	fmt.Println("El resultado de la suma variadica es", resultado)

	resultado2 := sumaVariadica(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15) //llamamos a la funcion variadica
	fmt.Println("El resultado de la suma variadica es", resultado2)

	imprimirDatos("Hola", 1, 1.2, true)

	resultado2 = factorial(5)
	fmt.Println("El factorial de 5 es", resultado2)

	// funciones anonimas
// una funcion anonima es una funcion que no tiene nombre
// se pueden asignar a variables y se pueden ejecutar en el momento
// se pueden pasar como parametros a otras funciones
// se pueden retornar como resultado de una funcion
// se pueden ejecutar en el momento
// se pueden ejecutar en una go routine

// en este caso, se va a hacer una funcion anonima que calcule el cuadrado de un numero
// y se va a ejecutar en el momento

	cuadrado := func(a int) int {
		return a * a
	}

	fmt.Println("El cuadrado de 5 es", cuadrado(5))

	saludo := func(name string){
		fmt.Println("Hola", name)
	}

	saludar("Leandro", saludo)

	// vamos a utilizar las funciones duplicar y triplicar
	fmt.Println("El doble de 5 es", aplicarFuncion(duplicar, 5))
	fmt.Println("El triple de 5 es", aplicarFuncion(triplicar, 5))

	// funciones de orden superior

	fmt.Println("El doble de 5 + uno es", double(addOne, 5))

	// funciones closures

	incremento := incrementar()
	fmt.Println(incremento())
} 

func suma(a, b int) int { //nombre de la funcion (parametros a pasar (maximos 3 es lo recomendado)) retorno de la funcion {cuerpo}
	return a + b
}

// funciones variadicas
func sumaVariadica(nums ...int) int { //se le pasa un numero indeterminado de parametros
	resultado := 0
	for _, num := range nums {
		resultado += num
	}
	return resultado
}

func imprimirDatos(datos ...interface{}){ //se le pasa un numero indeterminado de parametros
	for _, dato := range datos {
		fmt.Printf("%T - %v \n", dato, dato)
	}
}

// funciones recursivas
// una funcion recursiva es una funcion que se llama a si misma
// se debe tener un caso base para que no se llame infinitamente
// en este caso, se va a hacer una funcion que calcule el factorial de un numero

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func saludar(name string, f func(string)){
	f(name)
}

// otro ejemplo con duplicación y triplicación

func duplicar(a int) int {
	return a * 2
}

func triplicar(a int) int {
	return a * 3
}

func aplicarFuncion(f func(int) int, a int) int {
	return f(a)
}

// funciones de orden superior
// una funcion de orden superior es una funcion que recibe como parametro otra funcion
// o que retorna una funcion
// en este caso, se va a hacer una funcion que reciba una funcion y un numero
// y que retorne el resultado de aplicar la funcion al numero

func double(f func(int) int, a int) int {
	return f(a * 2)
}

func addOne(x int) int {
	return x + 1
}

// funciones closures
// un closure es una funcion que captura una variable del entorno en el que fue creada
// en este caso, se va a hacer una funcion que retorne otra funcion
// y que capture una variable del entorno

func incrementar() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}