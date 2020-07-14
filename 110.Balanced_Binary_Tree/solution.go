package main

import "fmt"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	hOfLeft := heightOfTree(root.Left, 0)
	hOfRight := heightOfTree(root.Right, 0)
	fmt.Println(hOfLeft, hOfRight)
	return abs(hOfLeft-hOfRight) < 2
}

func heightOfTree(root *TreeNode, h int) int {
	if root == nil {
		return h
	}

	return max(heightOfTree(root.Left, h+1), heightOfTree(root.Right, h+1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return a
	}
	return -a
}
