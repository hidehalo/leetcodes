package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pInLeft := inBST(root.Left, p.Val)
	qInLeft := inBST(root.Left, p.Val)
	if pInLeft && qInLeft {
		return lowestCommonAncestor(root.Left, p, q)
	} else if !pInLeft && !qInLeft && inBST(root.Right, p.Val) && inBST(root.Right, q.Val) {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func inBST(T *TreeNode, val int) bool {
	if T == nil {
		return false
	}
	if T.Val == val {
		return true
	}
	if T.Val > val {
		return inBST(T.Left, val)
	}
	return inBST(T.Right, val)
}
