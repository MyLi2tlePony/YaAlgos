package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if node.Left == nil && node.Right == nil {
		return value == node.Val
	}

	result := value - node.Val
	return hasPathSum(node.Left, result) || hasPathSum(node.Right, result)
}
