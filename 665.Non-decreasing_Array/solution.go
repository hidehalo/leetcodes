package main

func checkPossibility(nums []int) bool {
	nonDesc := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			if nonDesc >= 1 {
				return false
			}

			if i == 0 {
				nums[i] = nums[i+1] - 1
			} else if nums[i+1] >= nums[i-1] {
				nums[i] = nums[i-1]
			} else {
				nums[i+1] = nums[i]
			}
			nonDesc++
		}
	}

	return true
}
