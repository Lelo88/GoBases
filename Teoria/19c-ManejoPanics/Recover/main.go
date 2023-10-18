//recover intercepta el panic para que el cierre no sea abrupto

package main

import "fmt"

func isPair(num int) {
	defer func(){
		err := recover()
		if err!= nil {
            fmt.Println(err)
	}
}()

	if num%2!=0{
		panic("not a pair")
	}
	fmt.Println(num, "is a pair")		
}


func main() {
	num := 3

	isPair(num)

	fmt.Println("Funcion completada")	
}
