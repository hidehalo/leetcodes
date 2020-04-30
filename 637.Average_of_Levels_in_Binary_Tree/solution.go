package main

func averageOfLevels(root *TreeNode) []float64 {
	sumPerLevel := make([]float64, 0)
	childNumPerLevel := make([]int, 0)
	traversalSum(root, 0, &childNumPerLevel, &sumPerLevel)
	for i := 1; i < len(sumPerLevel); i++ {
		sumPerLevel[i] /= float64(childNumPerLevel[i])
	}
	return sumPerLevel
}

func traversalSum(node *TreeNode, level int, childNumPerLevel *[]int, sumPerLevel *[]float64) {
	if node != nil {
		if level >= len(*sumPerLevel) {
			(*sumPerLevel) = append((*sumPerLevel), float64(node.Val))
			(*childNumPerLevel) = append((*childNumPerLevel), 1)
		} else {
			(*sumPerLevel)[level] += float64(node.Val)
			(*childNumPerLevel)[level]++
		}
		traversalSum(node.Left, level+1, childNumPerLevel, sumPerLevel)
		traversalSum(node.Right, level+1, childNumPerLevel, sumPerLevel)
	}
}
