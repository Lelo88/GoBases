package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sol(t *testing.T) {
	
	var num1, num2, expected int64 = 2, 2, 4
	pnum1 := &num1
	pnum2 := &num2
	pexpected := &expected

	result, err := solution(pnum1, pnum2)

	require.NoError(t, err)
	require.Equal(t, pexpected, &result)
}