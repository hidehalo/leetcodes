package main

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func thirdMax(nums []int) int {
	max := make([]int, 4)
	for i := 0; i < 3; i++ {
		max[i] = MinInt
	}
	maxCount := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxCount = 1
			max[0] = nums[0]
		} else {
			if nums[i] == max[0] || nums[i] == max[1] || nums[i] == max[2] {
				continue
			}
			for j := 0; j < 3; j++ {
				if max[j] < nums[i] {
					if maxCount < 3 {
						maxCount++
					}
					// Insert
					for k := maxCount; k > j; k-- {
						max[k] = max[k-1]
					}
					max[j] = nums[i]
					break
				}
			}
		}
	}
	if maxCount < 3 {
		return max[0]
	}
	return max[maxCount-1]
}
