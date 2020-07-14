package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	hOfLeft, bOfLeft := heightOfTree(root.Left, 0)
	hOfRight, bOfRight := heightOfTree(root.Right, 0)
	if !bOfLeft || !bOfRight || abs(hOfLeft-hOfRight) >= 2 {
		return false
	}
	return true
}

func heightOfTree(root *TreeNode, h int) (int, bool) {
	if root == nil {
		return h, true
	}
	hOfLeft, bOfLeft := heightOfTree(root.Left, h+1)
	hOfRight, bOfRight := heightOfTree(root.Right, h+1)
	if !bOfLeft || !bOfRight || abs(hOfLeft-hOfRight) >= 2 {
		return max(hOfLeft, hOfRight), false
	}

	return max(hOfLeft, hOfRight), true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
