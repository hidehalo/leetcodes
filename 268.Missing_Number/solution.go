package main

func missingNumber(nums []int) int {
	excepted := 0
	for i := 1; i <= len(nums); i++ {
		excepted += i
	}

	actual := 0
	for i := 0; i < len(nums); i++ {
		actual += nums[i]
	}

	return excepted - actual
}
