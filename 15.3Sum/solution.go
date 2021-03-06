package main

import (
	"sort"
)

type T []int

func threeSum(nums T) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ret := make([][]int, 0, n)
	for i := 0; i < n-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			a := nums[i]
			start := i + 1
			end := n - 1
			for start < end {
				b := nums[start]
				c := nums[end]
				if a+b+c == 0 {
					ret = append(ret, []int{a, b, c})
					for start < end && b == nums[start+1] {
						start++
					}
					for start < end && c == nums[end-1] {
						end--
					}
					start++
					end--
				} else if a+b+c > 0 {
					end--
				} else {
					start++
				}
			}
		}
	}

	return ret
}
