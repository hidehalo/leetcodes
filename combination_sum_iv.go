package main

import (
	"fmt"
)

var dp []int

// Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.
func combinationSum4(nums []int, target int) int {
	dp = make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 1

	return reach(&nums, target)
}

func reach(nums *[]int, target int) int {
	if dp[target] != -1 {
		return dp[target]
	}
	cnt := 0
	for _, v := range *nums {
		if target >= v {
			cnt += reach(nums, target-v)
		}
	}
	dp[target] = cnt

	return cnt
}

func main() {
	arr := []int{1, 2, 3}
	target := 32
	ret := combinationSum4(arr, target)
	fmt.Println(ret)
}
