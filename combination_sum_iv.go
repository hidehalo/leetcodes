package main

import (
	"fmt"
)

// Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.
func combinationSum4(nums []int, target int) int {
	cnt := 0
	for _, v := range nums {
		next := target - v
		if next < 0 {
			continue
		}
		if next == 0 {
			cnt++
		}
		for _, tmp := range nums {
			if tmp <= next {
				cnt += combinationSum4(nums, next)
				break
			}
		}
	}

	return cnt
}

func main() {
	arr := []int{1, 2, 3}
	target := 32
	ret := combinationSum4(arr, target)
	fmt.Println(ret)
}
