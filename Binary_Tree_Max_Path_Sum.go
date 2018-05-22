func maxPathSum(root *TreeNode) int {
	return maxSum(root)
}

func maxSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftForestSum := maxSum(root.Left)
	rightForestSum := maxSum(root.Right)
	max := root.Val
	if root.Right != nil && rightForestSum > max {
		max = rightForestSum
	}
	if root.Left != nil && leftForestSum > max {
		max = leftForestSum
	}
	if root.Val+leftForestSum > max {
		max = root.Val + leftForestSum
	}
	if root.Val+rightForestSum > max {
		max = root.Val + rightForestSum
	}
	if root.Val+rightForestSum+leftForestSum > max {
		max = root.Val + rightForestSum + leftForestSum
	}

	return max
}