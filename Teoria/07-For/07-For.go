package main

import "fmt"

func main(){
	sum := 0

	fmt.Println("Sum vale: ", 0)

	for i:=0;i<=10;i++{ 
		sum+=i
	}

	fmt.Println("Sum ahora vale:",sum)
}