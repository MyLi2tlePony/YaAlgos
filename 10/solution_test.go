package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch(t *testing.T) {
	require.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
