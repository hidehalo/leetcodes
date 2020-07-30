package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	memo := make([]int, len(nums))
	memo[0] = nums[0]
	memo[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		memo[i] = max(nums[i]+memo[i-2], memo[i-1])
	}
	return memo[len(nums)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
