package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(node *TreeNode) []int {
	queue := list.New()
	queue.PushBack(node)

	var result []int
	for queue.Len() > 0 {
		treeNode := queue.Remove(queue.Front()).(*TreeNode)

		if treeNode != nil {
			result = append(result, treeNode.Val)
			queue.PushBack(treeNode.Left)
			queue.PushBack(treeNode.Right)
		}
	}

	return result
}
