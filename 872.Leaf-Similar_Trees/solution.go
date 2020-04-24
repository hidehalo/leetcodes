package main

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	seq1 := make([]int, 0)
	getLeafSeq(root1, &seq1)
	seq2 := make([]int, 0)
	getLeafSeq(root2, &seq2)
	if len(seq1) != len(seq2) {
		return false
	}
	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}

	return true
}

func getLeafSeq(root *TreeNode, seq *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*seq = append(*seq, root.Val)
	} else if root.Left != nil && root.Right != nil {
		getLeafSeq(root.Left, seq)
		getLeafSeq(root.Right, seq)
	} else if root.Left != nil {
		getLeafSeq(root.Left, seq)
	} else if root.Right != nil {
		getLeafSeq(root.Right, seq)
	}
}
