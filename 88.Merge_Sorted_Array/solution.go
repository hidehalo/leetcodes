package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			i++
		} else {
			nums1[i], nums2[j] = nums2[j], nums1[i]
			j++
		}
	}
	arr1 := nums2[0:j]
	arr2 := nums2[j:n]
	k := i
	i, j = 0, 0
	for k < m+n {
		if i >= len(arr1) {
			nums1[k] = arr2[j]
			j++
		} else if j >= len(arr2) {
			nums1[k] = arr1[i]
			i++
		} else {
			if arr1[i] <= arr2[j] {
				nums1[k] = arr1[i]
				i++
			} else {
				nums1[k] = arr2[j]
				j++
			}
		}
		k++
	}
}
