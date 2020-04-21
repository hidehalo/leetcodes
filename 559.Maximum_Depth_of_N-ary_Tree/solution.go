package main

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	max := 1
	for i := 0; i < len(root.Children); i++ {
		h := maxDepth(root.Children[i]) + 1
		if h > max {
			max = h
		}
	}

	return max
}
