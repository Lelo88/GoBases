package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_century(t *testing.T) {
	var year int64 = 2005
	pyear := &year

	result, err := century(pyear)

	require.NoError(t, err)
	require.Equal(t, int64(21), result)
}

func Test_centuryWithErr(t *testing.T) {
	var year int64 = 0
	pyear := &year

	result, err := century(pyear)

	require.Error(t, err)
	require.Equal(t, int64(0), result)
	require.ErrorContains(t, err, "no se puede realizar la operación")
}

func Test_main(t *testing.T) {
	main()
}
