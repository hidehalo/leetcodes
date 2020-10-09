package main

func findTilt(root *TreeNode) int {
	_, ret := sum(root)

	return ret
}

func sum(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	sumLeft, tiltLeft := sum(root.Left)
	sumRight, tiltRight := sum(root.Right)

	return root.Val + sumLeft + sumRight, abs(sumLeft-sumRight) + tiltLeft + tiltRight
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
