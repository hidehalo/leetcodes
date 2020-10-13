package main

func maxCount(m int, n int, ops [][]int) int {
	minM := m
	minN := n
	for _, op := range ops {
		if minM > op[0] {
			minM = op[0]
		}
		if minN > op[1] {
			minN = op[1]
		}
	}
	return minM * minN
}
