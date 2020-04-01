package main

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	rank := make(map[int]int)
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)
	for i := 0; i < len(copyNums); i++ {
		if _, ok := rank[copyNums[i]]; !ok {
			rank[copyNums[i]] = i
		}
	}
	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = rank[nums[i]]
	}

	return ret
}
