/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func search(root *TreeNode, ignore *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if k == root.Val && root != ignore {
		return true
	}
	if k < root.Val {
		return search(root.Left, ignore, k)
	}
	return search(root.Right, ignore, k)
}

func searchLoop(root *TreeNode, current *TreeNode, k int) bool {
	if current == nil {
		return false
	}
	if search(root, current, k-current.Val) {
		return true
	}
	return searchLoop(root, current.Left, k) ||
		searchLoop(root, current.Right, k)
}

func findTarget(root *TreeNode, k int) bool {
	return searchLoop(root, root, k)
}