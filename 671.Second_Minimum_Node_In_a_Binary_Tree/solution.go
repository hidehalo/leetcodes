package main

import (
	"sort"

	"../ds"
)

func findSecondMinimumValue(root *ds.TreeNode) int {
	arr := make([]int, 0)
	traversal(root, &arr)
	sort.Ints(arr)
	if len(arr) < 2 {
		return -1
	}
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > min {
			return arr[i]
		}
	}
	return -1
}

func traversal(root *ds.TreeNode, arr *[]int) {
	if root != nil {
		*arr = append(*arr, root.Val)
		traversal(root.Left, arr)
		traversal(root.Right, arr)
	}
}

// func minValue(root *ds.TreeNode) int {
// 	if root == nil {
// 		return -1
// 	} else if root.Left == nil && root.Right == nil {
// 		return root.Val
// 	} else if root.Left != nil && root.Right != nil {
// 		return min(root.Val, min(minValue(root.Left), minValue(root.Right)))
// 	} else if root.Left != nil {
// 		return min(root.Val, minValue(root.Left))
// 	}
// 	return min(root.Val, minValue(root.Right))
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
