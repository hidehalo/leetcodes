package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {

	} else if root.Left == nil && root.Right != nil {

	} else if root.Left != nil && root.Right == nil {

	} else {

	}
}

func countDepth(root *TreeNode, d int) int {
	if root.Left == nil && root.Right == nil {
		return d
	} else if root.Left == nil && root.Right != nil {
		return countDepth()
	} else if root.Left != nil && root.Right == nil {

	} else {

	}
}
