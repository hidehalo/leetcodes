package main

// Node is n-ary tree node's data structure
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, root.Val)
	for i := 0; i < len(root.Children); i++ {
		ret = append(ret, preorder(root.Children[i])...)
	}

	return ret
}
