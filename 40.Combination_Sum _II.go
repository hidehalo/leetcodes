package main

import (
	"fmt"
)

func combinationSum2(candidates []int, target int) [][]int {
	size := len(candidates)
	ret := make([][]int, 0, size)
	recur(candidates, target)

	return ret
}

func recur(candidates []int, target int) {
	size := len(candidates)
	fmt.Println(size, target)
	if size <= 0 {
		return
	}
	if target == candidates[0] {
		// fmt.Println(candidates[0], "fi")
	}
	if target < candidates[0] {
		// fmt.Println("no")
	}
	if target > candidates[0] {
		// fmt.Println(candidates[0])
	}
	for i := 0; i < size; i++ {
		recur(candidates[1:], target-candidates[0])
	}
}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	combinationSum2(candidates, target)
}
