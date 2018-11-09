package main

import (
	"fmt"
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

func main() {
	fmt.Println(threeSumClosest([]int{13, 2, 0, -14, -20, 19, 8, -5, -13, -3, 20, 15, 20, 5, 13, 14, -17, -7, 12, -6, 0, 20, -19, -1, -15, -2, 8, -2, -9, 13, 0, -3, -18, -9, -9, -19, 17, -14, -19, -4, -16, 2, 0, 9, 5, -7, -4, 20, 18, 9, 0, 12, -1, 10, -17, -11, 16, -13, -14, -3, 0, 2, -18, 2, 8, 20, -15, 3, -13, -12, -2, -19, 11, 11, -10, 1, 1, -10, -2, 12, 0, 17, -19, -7, 8, -19, -17, 5, -5, -10, 8, 0, -12, 4, 19, 2, 0, 12, 14, -9, 15, 7, 0, -16, -5, 16, -12, 0, 2, -16, 14, 18, 12, 13, 5, 0, 5, 6}, -59))
}
