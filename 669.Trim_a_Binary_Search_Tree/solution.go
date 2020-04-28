package main

func min(root *TreeNode) *TreeNode {
	p := root
	for p.Left != nil {
		p = p.Left
	}

	return p
}

func max(root *TreeNode) *TreeNode {
	p := root
	for p.Right != nil {
		p = p.Right
	}

	return p
}

func preccesor(root *TreeNode, v int) *TreeNode {
	p := root
	for p != nil {
		if p.Val == v {
			return p
		} else if p.Val < v {
			return max(p)
		}
		p = p.Left
	}

	return p
}

func succesor(root *TreeNode, v int) *TreeNode {
	p := root
	for p != nil {
		if p.Val == v {
			return p
		} else if p.Val > v {
			return min(p)
		}
		p = p.Right
	}

	return p
}

func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root.Val >= L && root.Val <= R {
		root.Left = preccesor(root, L)
		if root.Left != nil {
			if root.Left.Val == L {
				root.Left.Left = nil
			} else {
				root.Left = root.Left.Right
			}
		}
		root.Right = succesor(root, R)
		if root.Right != nil {
			if root.Right.Val == R {
				root.Right.Right = nil
			} else {
				root.Right = root.Right.Left
			}
		}
	} else if root.Val < L {
		root = succesor(root, L)
	} else {
		root = preccesor(root, R)
	}

	return root
}
