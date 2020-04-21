package main

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	bucket := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		bucket[arr1[i]]++
	}
	j := 0
	for i := 0; i < len(arr2); i++ {
		for bucket[arr2[i]] > 0 {
			arr1[j] = arr2[i]
			bucket[arr2[i]]--
			j++
		}
	}
	p := j
	for v, c := range bucket {
		if c > 0 {
			for i := 0; i < c; i++ {
				arr1[j] = v
				j++
			}
		}
	}
	sort.Ints(arr1[p:])

	return arr1
}
