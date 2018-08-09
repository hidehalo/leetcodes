package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	size := len(nums)
	tmp := helper(nums, 0, size-1, target)
	sizeT := len(tmp)
	low, high := -1, -1
	for i := 0; i < sizeT; i++ {
		if tmp[i] == target {
			low = i
			break
		}
	}
	for i := sizeT - 1; i >= low; i-- {
		if i < 0 {
			break
		}
		if tmp[i] == target {
			high = i
			break
		}
	}

	return []int{low, high}
}

func helper(nums []int, left int, right int, target int) []int {
	if left > right {
		return []int{}
	}
	mid := (left + right) / 2
	if nums[mid] > target {
		return helper(nums, left, mid-1, target)
	} else if nums[mid] < target {
		return helper(nums, mid+1, right, target)
	}

	return nums
}

func main() {
	fmt.Println(searchRange([]int{1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8}, 8))
}
