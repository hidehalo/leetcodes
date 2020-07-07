package main

func checkPossibility(nums []int) bool {
	nonDesc := 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] > nums[i] && nums[i] > nums[i+1] {
			return false
		}
		if nums[i] <= nums[i+1] && nums[i-1] <= nums[i] {
			continue
		} else {
			nonDesc++
			if nonDesc > 1 {
				return false
			}
		}

	}
	if len(nums) == 3 {
		return nums[1] <= nums[2]
	}
	return true
}
