package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	ret := binarySeach(nums, 0, len(nums)-1, target)

	return ret
}

func binarySeach(nums []int, left int, right int, target int) int {
	if left > right {
		return left
	}
	mid := (left + right) / 2
	if nums[mid] > target {
		return binarySeach(nums, left, mid-1, target)
	} else if nums[mid] < target {
		return binarySeach(nums, mid+1, right, target)
	}

	return mid
}

func main() {
	nums := []int{}
	target := 5
	fmt.Println(searchInsert(nums, target))
}
