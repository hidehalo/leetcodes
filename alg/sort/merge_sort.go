package main

import (
	"fmt"
)

func combine(l []int, r []int) []int {
	var i, j int
	a := make([]int, 0, len(l)+len(r))
	for i < len(l) || j < len(r) {
		if i >= len(l) {
			a = append(a, r[j])
			j++
		} else if j >= len(r) {
			a = append(a, l[i])
			i++
		} else if l[i] > r[j] {
			a = append(a, r[j])
			j++
		} else {
			a = append(a, l[i])
			i++
		}
	}

	return a
}

func mergeSort(a []int, l int, r int) []int {
	if l == r {
		return a[l : r+1]
	}
	mid := (l + r) / 2
	leftArr := mergeSort(a, l, mid)
	rightArr := mergeSort(a, mid+1, r)

	return combine(leftArr, rightArr)
}

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4}
	fmt.Println(mergeSort(nums, 0, len(nums)-1))
}
