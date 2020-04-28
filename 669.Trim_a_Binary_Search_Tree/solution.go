package main

func preccesor(root *TreeNode, v int) *TreeNode {
	p := root
	for p != nil {
		if p.Val == v {
			break
		} else if p.Val < v {
			break
		}
		p = p.Left
	}

	return p
}

func succesor(root *TreeNode, v int) *TreeNode {
	p := root
	for p != nil {
		if p.Val == v {
			break
		} else if p.Val > v {
			break
		}
		p = p.Right
	}

	return p
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root.Val >= L && root.Val <= R {
		root.Left = succesor(root.Left, L)
		if root.Left != nil {
			root.Left.Left = nil
		}
		root.Right = preccesor(root.Right, R)
		if root.Right != nil {
			root.Right.Right = nil
		}
	} else if root.Val < L {
		root = succesor(root.Right, L)
	} else {
		root = preccesor(root.Left, R)
	}

	return root
}
