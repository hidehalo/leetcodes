package main

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return float64(nums[0])
	} else if k >= len(nums) {
		s := 0
		for _, n := range nums {
			s += n
		}
		return float64(s) / float64(len(nums))
	}

	sums := make([]int, 0, len(nums)+1)
	sums = append(sums, 0)
	for i := 0; i < len(nums); i++ {
		sums = append(sums, sums[i]+nums[i])
	}
	maxSum := sums[k]
	for i := 0; i+k < len(sums); i++ {
		s := sums[i+k] - sums[i]
		if s > maxSum {
			maxSum = s
		}
	}
	return float64(maxSum) / float64(k)
}
