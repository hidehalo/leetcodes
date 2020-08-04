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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	} else if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	} else if root.Left != nil && root.Right != nil {
		return traversal(strconv.Itoa(root.Val), []*TreeNode{root.Left, root.Right})
	} else if root.Left != nil {
		return traversal(strconv.Itoa(root.Val), []*TreeNode{root.Left})
	}
	return traversal(strconv.Itoa(root.Val), []*TreeNode{root.Right})
}

func traversal(prefix string, childs []*TreeNode) []string {
	ret := make([]string, 0)
	for _, child := range childs {
		if child.Left == nil && child.Right == nil {
			ret = append(ret, prefix+"->"+strconv.Itoa(child.Val))
		} else {
			nextLevel := make([]*TreeNode, 0)
			if child.Left != nil {
				nextLevel = append(nextLevel, child.Left)
			}
			if child.Right != nil {
				nextLevel = append(nextLevel, child.Right)
			}
			ret = append(ret, traversal(prefix+"->"+strconv.Itoa(child.Val), nextLevel)...)
		}
	}
	return ret
}
