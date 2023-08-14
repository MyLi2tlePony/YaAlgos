package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSort(t *testing.T) {
	result := []int{3, 2, 5, 4, 1, 6, 7, 0}
	sort(result)
	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7}, result)

	result = []int{3, 2, 5, 4}
	sort(result)
	require.Equal(t, []int{2, 3, 4, 5}, result)

	result = []int{}
	sort(result)
	require.Equal(t, []int{}, result)

	result = []int{4, 3}
	sort(result)
	require.Equal(t, []int{3, 4}, result)
}
