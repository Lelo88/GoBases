package main

import (
	"errors"
	"fmt"
)

func solution(num1 *int64, num2 *int64) (int64, error) {
	if num1 == nil || num2 == nil {
		return 0, errors.New("no se puede realizar la operación")
	}

	return *num1 + *num2, nil
}

func main() {
	var num1, num2 int64 = 2, 2
	pnum1 := &num1
	pnum2 := &num2

	result, err := solution(pnum1, pnum2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
