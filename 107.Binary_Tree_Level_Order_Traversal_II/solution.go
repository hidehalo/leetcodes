package main

func levelOrderBottom(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	traversal(root, 0, &ret)
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}

	return ret
}

func traversal(root *TreeNode, h int, ret *[][]int) {
	if root != nil {
		if len((*ret)) <= h {
			(*ret) = append((*ret), make([]int, 0))
		}
		traversal(root.Left, h+1, ret)
		traversal(root.Right, h+1, ret)
		(*ret)[h] = append((*ret)[h], root.Val)
	}
}
