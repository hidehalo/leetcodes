package main

import (
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	min := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		j := i + 1
		if arr[j]-arr[i] < min {
			min = arr[j] - arr[i]
		}
	}
	ret := make([][]int, 0)
	for i := 0; i < len(arr)-1; i++ {
		j := i + 1
		if arr[j]-arr[i] == min {
			ret = append(ret, []int{arr[i], arr[j]})
		}
	}

	return ret
}
