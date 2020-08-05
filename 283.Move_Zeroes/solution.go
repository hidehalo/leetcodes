package main

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			j := i + 1
			for j < len(nums)-1 && nums[j] == 0 {
				j++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
