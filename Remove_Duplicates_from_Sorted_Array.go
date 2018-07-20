package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	i := 0
	for j := 1; j < size; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
