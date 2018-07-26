package main

import (
	"fmt"
)

func nextPermutation(nums []int) {
	size := len(nums)
	offset := -1
	for i := size - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			offset = i
			break
		}
	}
	if offset < 0 {
		reverse(nums)
	} else if offset <= size-2 {
		for i := size - 1; i > offset; i-- {
			if nums[i] > nums[offset] {
				swap(nums, i, offset)
				break
			}
		}
		reverse(nums[offset+1:])
	}
}

func swap(nums []int, i int, j int) {
	nums[i] ^= nums[j]
	nums[j] ^= nums[i]
	nums[i] ^= nums[j]
}

func reverse(nums []int) {
	i := 0
	j := len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
