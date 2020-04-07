package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(r *TreeNode, n int) {
	if r == nil {
		return
	}
	p := r
	var q *TreeNode
	for p != nil {
		q = p
		if p.Val < n {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	p = &TreeNode{
		n,
		nil,
		nil,
	}
	if q.Val <= n {
		q.Right = p
	} else {
		q.Left = p
	}
}

func newBST(nums []int) *TreeNode {
	r := &TreeNode{nums[0], nil, nil}
	for i := 1; i < len(nums); i++ {
		insert(r, nums[i])
	}
	return r
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	p := root
	for p != nil && (p.Val > R || p.Val < L) {
		if p.Val > R {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	if p == nil {
		return 0
	}

	return p.Val + rangeSumBST(p.Left, L, p.Val) + rangeSumBST(p.Right, p.Val, R)
}
