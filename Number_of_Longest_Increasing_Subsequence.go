package main

import (
	"fmt"
)

/**
 * Given an unsorted array of integers, find the number of longest increasing subsequence.
 */
func findNumberOfLIS(nums []int) int {
	x := 0
	size := len(nums)
	for i := 1; i < size-1; i++ {
		crt := nums[i]
		pre := nums[i-1]
		if pre <= crt {
			continue
		}
	}

	return x
}

func getLongest(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	count := 0
	guard := 0
	crt := nums[0]
	ret := make([]int, 0, size-1)
	for i := 0; i < size-1; i++ {
		next := nums[i+1]
		// fmt.Println(guard, crt, next, nums[i])
		if next < crt && next >= guard {
			crt = next
			fmt.Println(ret[len(ret)-1], crt)
			ret[len(ret)-1] = crt
		} else if next >= nums[i] {
			ret = append(ret, crt)
			guard = nums[i]
			crt = next
			count++
		}
	}
	fmt.Println(ret)

	return count
}

func main() {
	fmt.Println(getLongest([]int{1, 3, 5, 4, 7, 5, 6, 4, 8}))
}
