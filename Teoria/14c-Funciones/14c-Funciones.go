//ejemplo de clase en vivo de manejo de funciones a traves de orquestadores

package main

import "fmt"



const (
	OperadorSuma  = "+"
	OperadorResta = "-"
	OperadorMult  = "*"
	OperadorDiv   = "/"
)

type Operacion func(num1, num2 float64) float64 //tipo de variable declarada que sera una funcion que devolvera un flotante en especifico

func sumar(num1, num2 float64) (resultado float64) {
	resultado = num1 + num2	
	return
}

func restar(num1, num2 float64) (resultado float64) {
	resultado = num1 - num2
	return
}

func multiplicar(num1, num2 float64) (resultado float64) {
	resultado = num1 * num2
	return
}

func dividir(num1, num2 float64) (resultado float64) {

		if num2 == 0 {
			return 0 
		}
		resultado = num1 / num2
	
	return
}

func orquestarOperacion(nums []float64, operacion Operacion)(resultado float64){
	for index := range nums {
		if index == 0 {
			resultado = nums[index] //igualamos el resultado al primer indice en el caso de que metamos un cero implicitamente
			//podemos corroborar agregando un cero en la funcion de multiplicar o dividir
			continue
		}

		resultado = operacion(resultado, nums[index])
	}
	return 
}

func operacionAritmetica(operador string, nums ...float64) (resultado float64){
	switch operador{
	case OperadorSuma:
		resultado = orquestarOperacion(nums,sumar)
	case OperadorResta:
		resultado = orquestarOperacion(nums,restar)
	case OperadorMult:
		resultado = orquestarOperacion(nums,multiplicar)
	case OperadorDiv: 
		resultado = orquestarOperacion(nums,dividir)
	}
	
	return 
}

func operacionAritmetica2(operador string) (resultado Operacion){ //funcion que me retorna otra funcion. solo paso el operador y el resutlado sera una de las funciones de operacion aritmetica
	switch operador {
		case OperadorSuma:
			resultado = sumar
		case OperadorResta:
			resultado = restar
		case OperadorMult:
			resultado = multiplicar
		case OperadorDiv:
			resultado = dividir
	}
	return 
}

func main(){
	fmt.Println(operacionAritmetica(OperadorSuma, 1,2,3,4,5))
	fmt.Println(operacionAritmetica(OperadorResta, 1,2,3,4,5))
	fmt.Println(operacionAritmetica(OperadorMult, 1,2,3,4,5))
	fmt.Println(operacionAritmetica(OperadorDiv, 1,2,3,4,5))

	fmt.Printf("\n")

	//devolucion de funciones 
	oper := operacionAritmetica2(OperadorSuma) //llamo a la funcion pasando SOLO el string en este caso
	
	fmt.Println(oper(5.6,3.2)) //luego le paso los argumentos a la variable que he creado (es como si se transformara en una funcion)
}

// Los orquestadores son útiles cuando se necesita crear una función que depende de otra función
// y se quiere evitar repetir código
// un ejemplo de orquestador es la funcion operacionAritmetica2 que retorna una funcion.
