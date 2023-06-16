// Given a year, return the century it is in. The first century spans from the year 1 up to
// and including the year 100, the second from the year 101 up to and including the year 200, etc.

package main

import (
	"errors"
	"fmt"
)

func century (year *int64) (int64, error) {
	
	if *year <= 0 {
		return 0, errors.New("no se puede realizar la operación")
	}
	
	return (*year / 100) + 1, nil
}

func main() {
	var year int64 = 2005
	pyear := &year
	
	result, err := century(pyear)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(result)
}