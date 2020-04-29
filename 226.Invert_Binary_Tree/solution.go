package main

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		if root.Left != nil && root.Right != nil {
			invertTree(root.Left)
			invertTree(root.Right)
		} else if root.Left != nil && root.Right == nil {
			invertTree(root.Left)
		} else if root.Right != nil && root.Left == nil {
			invertTree(root.Right)
		}
	}

	return root
}
