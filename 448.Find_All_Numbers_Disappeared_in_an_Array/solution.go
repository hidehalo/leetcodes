package main

func findDisappearedNumbers(nums []int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	bucket := make([]int, len(nums))
	ret := make([]int, 0, len(nums)-1)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]-1]++
	}
	for i, v := range bucket {
		if v == 0 {
			ret = append(ret, i+1)
		}
	}

	return ret
}
