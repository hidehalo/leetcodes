package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	size := len(candidates)
	start, end := 0, size-1
	ret := make([][]int, 0, size)
	for start < end {
		if candidates[start]+candidates[end] > target {
			end--
		} else if candidates[start]+candidates[end] < target {
			start++
		} else {
			ret = append(ret, []int{start, target})
			for candidates[start] == candidates[start+1] && start < end {
				start++
			}
			for candidates[end] == candidates[end-1] && start < end {
				end--
			}
			start++
			end--
		}
	}
}

func main() {

}
