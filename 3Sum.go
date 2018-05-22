package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	size := len(nums)
	if size == 3 && (nums[0]+nums[1]+nums[2]) == 0 {
		return [][]int{{nums[0], nums[1], nums[2]}}
	}
	ret := make([][]int, 0, 1000)
	i := 0
	j := i + 1
	k := j + 1
	hashSet := make([]map[int]int, 0, size)
	for i < size-3 {
		if k > size-1 {
			j++
			if j > size-2 {
				i++
				j = i + 1
			}
			k = j + 1
		}
		var dup bool
		for _, set := range hashSet {
			tmp := make(map[int]int)
			tmp[nums[i]]++
			tmp[nums[j]]++
			tmp[nums[k]]++
			if tmp[nums[i]] == set[nums[i]] && tmp[nums[j]] == set[nums[j]] && tmp[nums[k]] == set[nums[k]] {
				dup = true
				break
			}
		}
		if dup {
			j++
			k = j + 1
			continue
		}
		if nums[i]+nums[j]+nums[k] == 0 {
			ret = append(ret, []int{nums[i], nums[j], nums[k]})
			set := make(map[int]int)
			set[nums[i]]++
			set[nums[j]]++
			set[nums[k]]++
			hashSet = append(hashSet, set)
		}
		k++
	}

	return ret
}

func main() {
	fmt.Println(threeSum([]int{-13, 5, 13, 12, -2, -11, -1, 12, -3, 0, -3, -7, -7, -5, -3, -15, -2, 14, 14, 13, 6, -11, -11, 5, -15, -14, 5, -5, -2, 0, 3, -8, -10, -7, 11, -5, -10, -5, -7, -6, 2, 5, 3, 2, 7, 7, 3, -10, -2, 2, -12, -11, -1, 14, 10, -9, -15, -8, -7, -9, 7, 3, -2, 5, 11, -13, -15, 8, -3, -7, -12, 7, 5, -2, -6, -3, -10, 4, 2, -5, 14, -3, -1, -10, -3, -14, -4, -3, -7, -4, 3, 8, 14, 9, -2, 10, 11, -10, -4, -15, -9, -1, -1, 3, 4, 1, 8, 1}))
}
