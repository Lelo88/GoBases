package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_productNumberAdjacent (t *testing.T) {
	var inputArray []int = []int{3, 6, -2, -5, 7, 3}

	productExpected := 21
	productResult := productNumbersAdjacent(inputArray)

	require.Equal(t, productExpected, productResult)
	require.NotNil(t, productResult)
}

func Test_productNumberAdjacent2 (t *testing.T) {
	var inputArray []int = []int{-1}

	productExpected := 0
	productResult := productNumbersAdjacent(inputArray)

	require.Equal(t, productExpected, productResult)
	require.NotNil(t, productResult)
}

func Test_main(t *testing.T) {
	main()
}