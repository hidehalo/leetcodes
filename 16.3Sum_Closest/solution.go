package main

import (
	"sort"
)

func distance(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	dmin := distance(target, nums[0]+nums[1]+nums[len(nums)-1])
	ret := nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			s := i + 1
			e := len(nums) - 1
			for s < e {
				tmp := nums[i] + nums[s] + nums[e]
				d := distance(target, tmp)
				// FIXME: pointer overstep
				// if d == dmin {
				// 	for s < e && nums[s] == nums[s+1] {
				// 		s++
				// 	}
				// 	for s < e && nums[e] == nums[e-1] {
				// 		e--
				// 	}
				// 	continue
				// }
				if tmp < target {
					s++
				} else {
					e--
				}
				if d < dmin {
					dmin = d
					ret = tmp
				}

			}
		}
	}

	return ret
}
