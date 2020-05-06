package main

// optimal: rotate only?
func increasingBST(root *TreeNode) *TreeNode {
	vals := traversal(root)
	head := &TreeNode{}
	p := head
	for i := 0; i < len(vals); i++ {
		node := &TreeNode{Val: vals[i]}
		p.Right = node
		p = p.Right
	}

	return head.Right
}

func traversal(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	left := traversal(node.Left)
	right := traversal(node.Right)
	vals := []int{}
	vals = append(vals, left...)
	vals = append(vals, node.Val)
	vals = append(vals, right...)

	return vals
}
