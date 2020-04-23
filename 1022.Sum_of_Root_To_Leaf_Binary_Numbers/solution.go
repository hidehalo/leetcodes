package main

func sumRootToLeaf(root *TreeNode) int {
	return procedure(root, 0)
}

func procedure(root *TreeNode, n int) int {
	if root == nil {
		return n
	}
	v := (n << 1) + root.Val

	return procedure(root.Left, v) + procedure(root.Right, v)

}
