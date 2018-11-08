package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	size := len(nums)
	ret := make([][]int, 0, size)
	if size < 4 {
		return ret
	}
	sort.Ints(nums)
	dp := make(map[int][][2]int)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			checked := true
			for _, pair := range dp[nums[i]+nums[j]] {
				if pair[0] == i || pair[1] == i || pair[0] == j || pair[1] == j {
					checked = false
					break
				}
			}
			if checked {
				dp[nums[i]+nums[j]] = append(dp[nums[i]+nums[j]], [2]int{i, j})
			}
		}
	}
	dup := make(map[int]bool)
	for t, p := range dp {
		// NOTE: result would be duplicated when t equals target and equals zero
		if target == 2*t {
			for i, pair := range p {
				for j := i + 1; j < len(p); j++ {
					ret = append(ret, []int{nums[p[j][0]], nums[p[j][1]], nums[pair[0]], nums[pair[1]]})
				}
			}
		} else {
			if v, ok := dup[t]; ok == true && v == true {
				continue
			}
			if pairsMirror, ok := dp[target-t]; ok == true && pairsMirror != nil {
				fmt.Println(t)
				for _, pair := range p {
					for _, mirror := range pairsMirror {
						ret = append(ret, []int{nums[mirror[0]], nums[mirror[1]], nums[pair[0]], nums[pair[1]]})
					}
				}
				dup[target-t] = true
			}
		}
		// start := 0
		// end := size - 1
		// for start < end-1 {
		// 	if target-t == nums[start]+nums[end] {
		// 		for start < end-1 && nums[start] == nums[start+1] {
		// 			start++
		// 		}
		// 		for start < end-1 && nums[end] == nums[end-1] {
		// 			end--
		// 		}
		// 		for _, pair := range p {
		// 			if pair[0] == start || pair[1] == start || pair[0] == end || pair[1] == end {
		// 				continue
		// 			}
		// 			ret = append(ret, []int{nums[pair[0]], nums[pair[1]], nums[start], nums[end]})
		// 		}
		// 	}
		// 	start++
		// 	end--
		// }
	}

	return ret
}

func main() {
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
