package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Printf("%#v", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("First argument:", os.Args[1])
	fmt.Println("Second argument:", os.Args[2])
	fmt.Println("Third argument:", os.Args[3])	

	fmt.Println("Number of arguments:", len(os.Args))
	fmt.Println("Hola")
}

/*
 % go build -o greeter           
% ./greeter hi leandro villalba
[]string{"./greeter", "hi", "leandro", "villalba"}Path: ./greeter
First argument: hi
Second argument: leandro
Third argument: villalba
Number of arguments: 4
*/

