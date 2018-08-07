package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)
	stopIndex := (size1 + size2) / 2
	if size1 <= 0 {
		return float64(nums2[stopIndex])
	} else if size2 <= 0 {
		return float64(nums1[stopIndex])
	}
	var isEven bool
	if (size1+size2)%2 == 0 {
		isEven = true
	}
	p1 := 0
	p2 := 0
	for p1 < size1 && p2 < size2 {
		if p1+p2 >= stopIndex {
			break
		}
		if nums1[p1] > nums2[p2] {
			p2++
		} else {
			p1++
		}
	}
	if isEven == true {
		if p2 >= size2 {
			p2 = size2 - 1
		}
		if p1 >= size1 {
			p1 = size1 - 1
		}

		return 0.5 * float64(nums1[p1]+nums2[p2])
	}
	if nums1[p1] < nums2[p2] {
		return float64(nums1[p1])
	}

	return float64(nums2[p2])
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
