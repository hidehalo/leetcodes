package main

import (
	"fmt"
)

func findMaximumXOR(nums []int) int {
	size := len(nums)
	var max, xorMax int
	max = nums[0]
	xorMax = 0
	for i := 1; i < size; i++ {
		if nums[i] > max {
			max = nums[i]
			continue
		}
	}
	for i := 0; i < size; i++ {
		if nums[i]^max > xorMax {
			xorMax = nums[i] ^ max
		}
	}

	return xorMax
}

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}
