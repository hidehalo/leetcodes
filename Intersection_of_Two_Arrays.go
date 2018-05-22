package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	size1 := len(nums1)
	size2 := len(nums2)
	if size1 <= 0 || size2 <= 0 {
		return []int{}
	}
	hashMap := make(map[int]bool)
	for _, vi := range nums1 {
		hashMap[vi] = true
	}
	ret := make([]int, 0, size1+size2)
	for _, vj := range nums2 {
		if hashMap[vj] == true {
			hashMap[vj] = false
			ret = append(ret, vj)
		}
	}

	return ret
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
