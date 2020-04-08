package main

func combinationSum2(candidates []int, target int) [][]int {
	size := len(candidates)
	ret := make([][]int, 0, size)
	for i := 0; i < size-1; i++ {
		recur(candidates[i:], 0, target)
	}

	return ret
}

func recur(candidates []int, offset int, target int) {
	size := len(candidates)
	if offset >= size {
		return
	}
	if target == 0 {
		return
	}
	if target < 0 {
		return
	}
	i := offset
	for ; i < size-1; i++ {
		if target < candidates[i] {
			continue
		} else {
			break
		}
	}
	recur(candidates, i+1, target-candidates[i])
}
