package main

func checkPossibility(nums []int) bool {
	nonDesc := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nonDesc++
			if nonDesc > 1 {
				return false
			}
		}
	}
	return true
}
