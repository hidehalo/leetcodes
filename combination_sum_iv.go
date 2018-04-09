package main

import (
	"fmt"
)

func combinationSum4(nums []int, target int) int {
	cnt := 0
	for _, v := range nums {
		next := target - v
		if next < 0 {
			continue
		}
		if next == 0 {
			return cnt + 1
		}
		for _, tmp := range nums {
			if tmp == next {
				cnt += combinationSum4(nums, next)
				break
			}
		}
	}

	return cnt
}

func main() {
	arr := []int{2, 1, 3}
	target := 4
	ret := combinationSum4(arr, target)
	fmt.Println(ret)
}
