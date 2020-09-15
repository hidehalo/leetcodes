package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	tmp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			tmp = 0
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
