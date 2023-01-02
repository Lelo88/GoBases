package main

import "fmt"

func main(){

	//operadores aritmeticos
	num1:=10
	num2:=5

	fmt.Println("Operadores aritmeticos")
	
	fmt.Println(num1 + num2)
	fmt.Println(num2-num1)
	fmt.Println(num2*num1)
	fmt.Println(num2/num1)
	fmt.Println((float64(num2))/float64(num1)) //como el lenguaje es muy estricto, 
	//tengo que castear las dos variables para que me de un resultado decimal
	fmt.Println(num1%num2)

	num1 ++
	fmt.Println("Ahora num1 vale:", num1)
	num1 --
	fmt.Println("Ahora num1 vale:", num1)

	//operadores relacionales
	fmt.Println("Operadores relacionales")
	fmt.Println(num1==num2) //false
	fmt.Println(num1>num2)  //true
	fmt.Println(num1<num2)  //false
	fmt.Println(num1<=num2) //false
	fmt.Println(num1>=num2) //true
	fmt.Println(num1!=num2) //true

	//operadores lógicos
	fmt.Println("Operadores lógicos")
	a := true
	b := false

	fmt.Println(a && b) 	//false
	fmt.Println(a || b) 	//true 
	fmt.Println(!(a && b))  //true

	//operadores de asignación
	fmt.Println("Operadores de asignación")
	num1+=2 //no funcionan dentro de las funciones. Utilizadas mayormente en ciclos
	fmt.Println("num1(10) += 2 = ", num1)
	num1-=10 
	fmt.Println("num1(12) -= 10 = ", num1)
	num1*=10
	fmt.Println("num1(2) *= 10 = ", num1)
	num1/=5
	fmt.Println("num1(20) /= 10 = ", num1)

	
}	