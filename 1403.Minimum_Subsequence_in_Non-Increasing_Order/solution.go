package main

import "sort"

func minSubsequence(nums []int) []int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	subsum := 0
	ret := make([]int, 0)
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		ret = append(ret, nums[i])
		if sum-nums[i] < subsum+nums[i] {
			break
		}
		subsum += nums[i]
		sum -= nums[i]
	}

	return ret
}
