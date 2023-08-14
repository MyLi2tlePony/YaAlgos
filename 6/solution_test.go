package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculateSum(t *testing.T) {
	root := &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val:   5,
			Right: nil,
			Left:  nil,
		},
		Left: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val:   3,
				Right: nil,
				Left:  nil,
			},
			Left: &TreeNode{
				Val:   4,
				Right: nil,
				Left:  nil,
			},
		},
	}

	require.Equal(t, []int{5, 5, 5, 4, 3}, bfs(root))
}
