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

	require.True(t, BranchValue(root, 10))
	require.True(t, BranchValue(root, 13))
	require.True(t, BranchValue(root, 14))
	require.False(t, BranchValue(root, 9))
	require.False(t, BranchValue(root, 100))

}
