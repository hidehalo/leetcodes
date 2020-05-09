package main

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}
	total := len(nums) * len(nums[0])
	if r*c > total || total/r > c || total/c > r {
		return nums
	}
	ret := make([][]int, r)
	ri, ci := 0, 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			ret[ri] = append(ret[ri], nums[i][j])
			ci = (ci + 1) % c
			if ci == 0 {
				ri++
			}
		}
	}

	return ret
}
