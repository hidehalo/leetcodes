package main

import (
	"fmt"
)

func insertSort(a []int) {
	size := len(a)
	for i := 1; i < size; i++ {
		key := a[i]
		pos := i
		for j := i - 1; j >= 0; j-- {
			if key < a[j] {
				a[j+1] = a[j]
				pos = j
			}
		}
		a[pos] = key
	}
}

func main() {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4}
	insertSort(nums)
	fmt.Println(nums)
}
