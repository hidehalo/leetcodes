package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{
		Left:  nil,
		Right: nil,
		Val:   nums[mid],
	}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if mid+1 <= len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}
