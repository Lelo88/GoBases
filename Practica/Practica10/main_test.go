package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solution(t *testing.T) {

	var num1, num2, expected int64 = 2, 2, 4
	pnum1 := &num1
	pnum2 := &num2
	pexpected := &expected

	result, err := solution(pnum1, pnum2)

	require.NoError(t, err)
	require.Equal(t, pexpected, &result)
}

func Test_solutionWithErr(t *testing.T) {

	var num1, num2 int64 = 0, 0
	pnum1 := &num1
	pnum2 := &num2

	result, err := solution(pnum1, pnum2)

	require.Error(t, err)
	require.Equal(t, int64(0), result)
	require.ErrorContains(t, err, "no se puede realizar la operación")
}

func Test_main(t *testing.T) {
	main()
}
