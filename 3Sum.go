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
	hashSet := make(map[int]bool)
	for i < size-3 {
		if k > size-1 {
			j++
			if j > size-2 {
				i++
				j = i + 1
			}
			k = j + 1
		}
		fmt.Println("nums:", nums[i], nums[j], nums[k])
		fmt.Println("hash set:", hashSet[nums[i]], hashSet[nums[j]], hashSet[nums[k]])
		if hashSet[nums[i]] && hashSet[nums[j]] && hashSet[nums[k]] {
			j++
			k = j + 1
			continue
		}
		// fmt.Println("index:", i, j, k)
		if nums[i]+nums[j]+nums[k] == 0 {
			if !hashSet[nums[i]] || !hashSet[nums[j]] || !hashSet[nums[k]] {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				hashSet[nums[i]] = true
				hashSet[nums[j]] = true
				hashSet[nums[k]] = true
			}
		}
		k++
	}

	return ret
}

func main() {
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}))
}
