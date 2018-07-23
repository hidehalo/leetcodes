package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	i := 0
	for j := 0; j < size; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

func main() {
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(removeElement(arr, val))
	fmt.Println(arr)
}
