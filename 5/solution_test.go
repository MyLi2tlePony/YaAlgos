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

	require.True(t, hasPathSum(root, 10))
	require.True(t, hasPathSum(root, 13))
	require.True(t, hasPathSum(root, 14))
	require.False(t, hasPathSum(root, 9))
	require.False(t, hasPathSum(root, 100))

}
