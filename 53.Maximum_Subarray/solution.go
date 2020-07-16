package main

func maxSubArray(nums []int) int {
	ret := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ret = max(sum, ret)
		sum = max(sum, 0)
	}
	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
