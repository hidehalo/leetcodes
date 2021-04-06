package main

import "strconv"

func itoa(n int) string {
	return strconv.Itoa(n)
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	} else if len(nums) == 1 {
		return []string{itoa(nums[0])}
	}
	ret := make([]string, 0, len(nums))
	start := nums[0]
	rangeLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] || nums[i]-nums[i-1] == 1 {
			rangeLen++
			if i == len(nums)-1 {
				if rangeLen > 1 {
					ret = append(ret, itoa(start)+"->"+itoa(nums[i]))
				} else {
					ret = append(ret, itoa(start))
				}
				rangeLen = 0
			}
			continue
		}
		if rangeLen > 1 {
			ret = append(ret, itoa(start)+"->"+itoa(nums[i-1]))
		} else {
			ret = append(ret, itoa(start))
		}
		start = nums[i]
		rangeLen = 1
	}
	if rangeLen > 0 {
		ret = append(ret, itoa(start))
	}

	return ret
}
