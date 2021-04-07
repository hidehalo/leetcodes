package main

import "sort"

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		max := 1
		for _, v := range nums {
			max *= v
		}
		return max
	}
	sort.Ints(nums)
	endOfArrIdx := len(nums) - 1
	max1 := nums[endOfArrIdx] * nums[0] * nums[1]
	max2 := nums[endOfArrIdx] * nums[endOfArrIdx-1] * nums[endOfArrIdx-2]
	if max1 > max2 {
		return max1
	}
	return max2
}
