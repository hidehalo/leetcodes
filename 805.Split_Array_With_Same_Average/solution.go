package main

func splitArraySameAverage(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	// TODO: impl&test
}

func hasKSum(A []int, k, target int) bool {
	if target < 0 {
		return false
	} else if target == 0 && k == 0 {
		return true
	}
	for i := 0; i < len(A); i++ {
		if hasKSum(A[i:], k-1, target-A[i]) {
			return true
		}
	}
	return false
}
