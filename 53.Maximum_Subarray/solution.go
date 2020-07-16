package main

func maxSubArray(nums []int) int {
	mid := len(nums) / 2

}

func arraySum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
