package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	height(root)
	return diameter
}

func height(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	leftHeight := height(tree.Left)
	rightHeight := height(tree.Right)
	if leftHeight+rightHeight > diameter {
		diameter = leftHeight + rightHeight
	}
	if leftHeight >= rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}
