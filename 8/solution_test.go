package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindArrWithMaxSum(t *testing.T) {
	require.Equal(t, 6, FindArrWithMaxSum([]int{2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
