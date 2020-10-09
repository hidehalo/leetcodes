package main

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	} else if root.Left != nil {
		return root.Left.Val + findTilt(root.Left)
	} else if root.Right != nil {
		return root.Right.Val + findTilt(root.Right)
	}
	return abs(root.Right.Val-root.Left.Val) + findTilt(root.Right) + findTilt(root.Left)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
