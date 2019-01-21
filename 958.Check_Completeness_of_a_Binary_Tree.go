package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	} else {
		return false
	}
	return isCompleteTree(root.Left) && isCompleteTree(root.Right)
}

func NewTree(nodes []int) *TreeNode {

}

func main() {

}
