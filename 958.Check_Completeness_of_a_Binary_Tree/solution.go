package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	treeList := make([]*TreeNode, 0, 1000)
	treeList = append(treeList, root)
	i := 0
	for i < len(treeList) && treeList[i] != nil {
		treeList = append(treeList, treeList[i].Left)
		treeList = append(treeList, treeList[i].Right)
		i++
	}
	for i < len(treeList) && treeList[i] == nil {
		i++
	}

	return i == len(treeList)
}

func NewTree(nodes []int) *TreeNode {
	treeList := make([]*TreeNode, 0, 1000)
	for i := 0; i < len(nodes); i++ {
		if nodes[i] == -1 {
			treeList = append(treeList, nil)
		} else {
			treeList = append(treeList, &TreeNode{nodes[i], nil, nil})
		}
	}
	for i := 0; i < len(treeList); i++ {
		l := 2*i + 1
		r := l + 1
		if l < len(treeList) {
			treeList[i].Left = treeList[l]
		}
		if r < len(treeList) {
			treeList[i].Right = treeList[r]
		}
	}

	return treeList[0]
}
