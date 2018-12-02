package main

import (
	"fmt"
)

func combinationSum2(candidates []int, target int) [][]int {
	size := len(candidates)
	ret := make([][]int, 0, size)
	recur(&candidates, 0, target)

	return ret
}

func recur(candidates *[]int, offset int, target int) {
	size := len((*candidates))
	if size <= 0 {
		return
	}
	if target == 0 {
		fmt.Println((*candidates)[0], "fi")
		return
	}
	if target < 0 {
		return
	}
	i := offset
	for ; i < size-1; i++ {
		if target < (*candidates)[i] {
			continue
		} else {
			break
		}
	}
	recur(candidates, i+1, target-(*candidates)[i])
}

func main() {
	candidates := []int{2, 5, 2, 1, 2}
	target := 5
	combinationSum2(candidates, target)
}
