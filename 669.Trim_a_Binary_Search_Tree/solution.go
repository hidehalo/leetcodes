package main

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < L {
		root.Left = nil
		return trimBST(root.Right, L, R)
	} else if root.Val > R {
		root.Right = nil
		return trimBST(root.Left, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)

	return root
}
