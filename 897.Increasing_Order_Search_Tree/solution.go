package main

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	} else if root.Left != nil && root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		root.Right = increasingBST(root.Right)
	} else if root.Left == nil && root.Right != nil {
		root.Right = increasingBST(root.Right)
	} else {
		tmp := root.Right
		root.Right = increasingBST(root.Left)
		root.Left = nil
		p := root
		for p.Right != nil {
			p = p.Right
		}
		p.Right = increasingBST(tmp)
	}

	return root
}
