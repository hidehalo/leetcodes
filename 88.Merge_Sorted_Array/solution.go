package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n <= 0 {
		return
	} else if m <= 0 {
		nums1 = append(nums1, nums2...)
	} else {
		i, j := 0, 0
		for i < m+n {
			if i >= m {
				nums1[i] = nums2[i-m]
			} else {
				if nums1[i] > nums2[j] {
					nums1[i], nums2[j] = nums2[j], nums1[i]
					if j < n-1 && nums2[j] > nums2[j+1] { // re-order nums2 if necessary
						sort.Ints(nums2[j:n]) // keep nums2 ordered
					} else { // move on
						j++
					}
				}
			}
			i++
		}
	}
}
