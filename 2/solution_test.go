package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSort(t *testing.T) {
	result := []int{3, 2, 4, 1}
	sort(result)
	require.Equal(t, []int{1, 2, 3, 4}, result)

	result = []int{}
	sort(result)
	require.Equal(t, []int{}, result)

	result = []int{4, 3}
	sort(result)
	require.Equal(t, []int{3, 4}, result)
}
