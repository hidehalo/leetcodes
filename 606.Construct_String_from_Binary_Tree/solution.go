package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func helper(ret *[]byte, t *TreeNode) {
	if t != nil {
		*ret = append(*ret, strconv.Itoa(t.Val)...)
		if t.Left != nil && t.Right != nil {
			*ret = append(*ret, '(')
			helper(ret, t.Left)
			*ret = append(*ret, ")("...)
			helper(ret, t.Right)
			*ret = append(*ret, ')')
		} else if t.Left != nil {
			*ret = append(*ret, '(')
			helper(ret, t.Left)
			*ret = append(*ret, ')')
		} else if t.Right != nil {
			*ret = append(*ret, "()("...)
			helper(ret, t.Right)
			*ret = append(*ret, ')')
		}
	}
}

func tree2str(t *TreeNode) string {
	ret := make([]byte, 0)
	helper(&ret, t)
	return string(ret)
}
