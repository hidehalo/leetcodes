package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums2KeyMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		nums2KeyMap[nums2[i]] = i + 1
	}
	ret := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if nums2KeyMap[nums1[i]] == 0 {
			ret = append(ret, -1)
		} else {
			find := 0
			for j := nums2KeyMap[nums1[i]] - 1; j < len(nums2); j++ {
				if nums2[j] > nums1[i] {
					ret = append(ret, nums2[j])
					find = 1
					break
				}
			}
			if find == 0 {
				ret = append(ret, -1)
			}
		}
	}

	return ret
}
