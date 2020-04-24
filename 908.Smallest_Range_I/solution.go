package main

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for i := 0; i < len(A); i++ {
		if A[i] < min {
			min = A[i]
		}
		if A[i] > max {
			max = A[i]
		}
	}
	ret := max - min - 2*K
	if ret > 0 {
		return ret
	}
	return 0
}
