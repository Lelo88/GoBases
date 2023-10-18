package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_matrixElementSum(t *testing.T) {
	var matrix [][]int = [][]int{
		{0, 1, 1, 2}, 
		{0, 5, 0, 0}, 
		{2, 0, 3, 3},
	}

	sumExpected := 9
	sumResult := matrixElementsSum(matrix)

	require.Equal(t, sumExpected, sumResult)
	require.NotNil(t, sumResult)
}

func Test_main(t *testing.T) {
	main()
}