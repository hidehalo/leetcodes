package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// result build for LC bug
type result struct {
	store []int
}

func findMode(root *TreeNode) []int {
	result := &result{make([]int, 0)}
	traversal(root, result)
	if len(result.store) <= 1 {
		return result.store
	}
	sort.Ints(result.store)
	fmt.Println(result.store)
	groupCount := 1
	groupVal := result.store[0]
	result.store = append(result.store, result.store[len(result.store)-1]+1)
	maxCount := 0
	ret := []int{groupVal}
	for i := 1; i < len(result.store); i++ {
		if result.store[i] == groupVal {
			groupCount++
			continue
		}
		fmt.Println("maxCount=", maxCount, "groupCount=", groupCount)
		if maxCount < groupCount {
			ret = []int{groupVal}
			maxCount = groupCount
		} else if maxCount == groupCount {
			ret = append(ret, groupVal)
		}
		groupVal = result.store[i]
		groupCount = 1
	}
	return ret
}

func traversal(root *TreeNode, ret *result) {
	if root == nil {
		return
	}
	if root.Right != nil && root.Left != nil {
		ret.store = append(ret.store, root.Val)
		traversal(root.Left, ret)
		traversal(root.Right, ret)

	} else if root.Left != nil {
		ret.store = append(ret.store, root.Val)
		traversal(root.Left, ret)
	} else if root.Right != nil {
		ret.store = append(ret.store, root.Val)
		traversal(root.Right, ret)
	} else {
		ret.store = append(ret.store, root.Val)
	}
}
