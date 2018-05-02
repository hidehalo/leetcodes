package main

import (
	"fmt"
)

func findMaximumXOR(nums []int) int {
	size := len(nums)
	var flag, max, xorMax int
	max = nums[0]
	flag = 0
	xorMax = 0
	for i := 1; i < size; i++ {
		if nums[i] > max {
			max = nums[i]
			continue
		}
		if nums[i]&max < flag {
			flag = nums[i] & max
			xorMax = max ^ flag
		}
	}

	return xorMax
}

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}
