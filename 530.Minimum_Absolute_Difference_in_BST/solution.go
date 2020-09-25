package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Optimized to T=O(n),S=O(1) ???
func getMinimumDifference(root *TreeNode) int {
	mins := make([]int, 0)
	travesal(root, &mins)
	min := mins[len(mins)-1]
	for i := 0; i < len(mins)-1; i++ {
		dis := abs(mins[i] - mins[i+1])
		if dis == 0 {
			return 0
		}
		if dis < min {
			min = dis
		}
	}
	return min
}

func travesal(root *TreeNode, mins *[]int) {
	if root == nil {
		return
	}
	travesal(root.Left, mins)
	*mins = append(*mins, root.Val)
	travesal(root.Right, mins)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
