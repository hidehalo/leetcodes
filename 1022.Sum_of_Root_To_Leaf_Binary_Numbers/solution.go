package main

func sumRootToLeaf(root *TreeNode) int {
	return procedure(root, 0)
}

func procedure(root *TreeNode, n int) int {
	if root == nil {
		return 0
	}
	v := (n << 1) + root.Val
	if root.Left == nil && root.Right == nil {
		return v
	}
	if root.Left != nil && root.Right != nil {
		return procedure(root.Left, v) + procedure(root.Right, v)
	} else if root.Left != nil {
		return procedure(root.Left, v)
	}
	return procedure(root.Right, v)
}
