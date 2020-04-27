package main

import (
	"fmt"

	"../ds"
)

func increasingBST(root *ds.TreeNode) *ds.TreeNode {
	p := root
	var q *ds.TreeNode
	for p.Left != nil {
		q = p
		p = p.Left
	}
	root = p
	procedure(q)

	return root
}
func procedure(root *ds.TreeNode) {
	if root.Left != nil {
		fmt.Printf("proceduring %v\n", root)
		rightRotate(root)
		procedure(root.Right)
	}
}

// func leftRotate(p *ds.TreeNode) {
// 	hold := p.Right
// 	p.Right.Left = p
// 	p.Right = hold.Left
// 	fmt.Printf("p=%v\tp.Left=%v\tp.Right=%v\n", p, p.Left, p.Right)

// }

func rightRotate(p *ds.TreeNode) {
	hold := p.Left.Right
	p.Left.Right = p
	p.Left = hold
	fmt.Printf("p=%v\tp.Left=%v\tp.Right=%v\n", p, p.Left, p.Right)
}
