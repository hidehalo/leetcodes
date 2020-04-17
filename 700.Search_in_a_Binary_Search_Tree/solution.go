package main

import "../ds"

func searchBST(root *ds.TreeNode, val int) *ds.TreeNode {
	p := root
	for p != nil {
		if p.Val > val {
			p = p.Left
		} else if p.Val < val {
			p = p.Right
		} else {
			break
		}
	}

	return p
}
