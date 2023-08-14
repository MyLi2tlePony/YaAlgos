package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateSum(t *testing.T) {
	root := &Node{
		Value: 5,
		R: &Node{
			Value: 5,
			R:     nil,
			L:     nil,
		},
		L: &Node{
			Value: 5,
			R: &Node{
				Value: 3,
				R:     nil,
				L:     nil,
			},
			L: &Node{
				Value: 4,
				R:     nil,
				L:     nil,
			},
		},
	}

	require.Equal(t, 22, calculateSum(root))
}
