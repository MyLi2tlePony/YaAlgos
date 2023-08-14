package main

type Node struct {
	Value int
	R     *Node
	L     *Node
}

func BranchValue(node *Node, value int) bool {
	if node == nil {
		return value == 0
	}

	result := value - node.Value
	return BranchValue(node.L, result) || BranchValue(node.R, result)
}
