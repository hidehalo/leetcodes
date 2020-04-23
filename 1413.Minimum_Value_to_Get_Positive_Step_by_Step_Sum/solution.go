package main

func minStartValue(nums []int) int {
	ret := 1
	for i := 1; ; i++ {
		sum := i
		tested := 1
		for j := 0; j < len(nums); j++ {
			sum += nums[j]
			if sum < 1 {
				tested = 0
				break
			}
		}
		if tested == 1 {
			ret = i
			break
		}
	}

	return ret
}
