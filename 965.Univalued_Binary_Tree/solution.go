package main

func isUnivalTree(root *TreeNode) bool {
    if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		return root.Val == root.Left.Val && 
			root.Val == root.Right.Val && 
			isUnivalTree(root.Right) && 
			isUnivalTree(root.Left)
	} else if root.Left != nil {
		return root.Val == root.Left.Val && 
			isUnivalTree(root.Left)
	} else if root.Right != nil {
		return root.Val == root.Right.Val &&
			isUnivalTree(root.Right)
	}

	return true
}
