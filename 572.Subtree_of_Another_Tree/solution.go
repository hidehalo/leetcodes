package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	sb := treeToInts(s)
	tb := treeToInts(t)
	if len(tb) > len(sb) {
		return false
	}
	for i := 0; i < len(sb); i++ {
		if sb[i] == tb[0] && len(sb)-i >= len(tb) {
			eq := true
			for j := 1; j < len(tb); j++ {
				if sb[i+j] != tb[j] {
					eq = false
					break
				}
			}
			if eq {
				return true
			}
		}
	}
	return false
}

func treeToInts(root *TreeNode) []int {
	ints := []int{}
	if root != nil {
		ints = append(ints, root.Val)
		ints = append(ints, treeToInts(root.Left)...)
		ints = append(ints, treeToInts(root.Right)...)
	} else {
		ints = append(ints, -65536)
	}
	return ints
}
