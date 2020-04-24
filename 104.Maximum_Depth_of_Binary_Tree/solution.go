package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left != nil && root.Right != nil {
		return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	} else if root.Right != nil {
		return 1 + maxDepth(root.Right)
	}
	return 1 + maxDepth(root.Left)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
