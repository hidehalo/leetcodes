package main

// Node is n-ary tree node's data structure
type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	for i := 0; i < len(root.Children); i++ {
		ret = append(ret, postorder(root.Children[i])...)
	}
	ret = append(ret, root.Val)

	return ret
}
