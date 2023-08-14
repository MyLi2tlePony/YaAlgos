package main

type Node struct {
	Value int
	R     *Node
	L     *Node
}

func calculateSum(node *Node) int {
	if node == nil {
		return 0
	}

	return calculateSum(node.R) + calculateSum(node.L) + node.Value
}
