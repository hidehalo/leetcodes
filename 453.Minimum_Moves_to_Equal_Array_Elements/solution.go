package main

// Assume need to move `m` times, and the final elements are all `x`, then:
// sum(nums) + m * (len(nums) - 1) = x * len(nums)
// m = x - min(nums)
// => x = m + min(nums)
// => m = sum(nums) - min(nums) * len(nums)
func minMoves(nums []int) int {
	sum := nums[0]
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - min*len(nums)
}
