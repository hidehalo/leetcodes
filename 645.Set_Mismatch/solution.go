package main

func findErrorNums(nums []int) []int {
	dup := make([]int, len(nums))
	dupVal := 0
	sumObtaind := 0
	for _, n := range nums {
		dup[n-1]++
		sumObtaind += n
		if dup[n-1] > 1 {
			dupVal = n
		}
	}
	sumWanted := (1 + len(nums)) * len(nums) / 2
	return []int{dupVal, sumWanted + dupVal - sumObtaind}
}
