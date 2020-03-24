package main

import "sort"

func splitArraySameAverage(A []int) bool {
	sum := 0
	n := len(A)
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for k := 1; k <= n/2; k++ {
		// fmt.Printf("%d*%d%%%d=%d\n", sum, k, n, sum*k%n)
		if sum*k%n == 0 && hasKSum(A, k, sum*k/n) {
			return true
		}
	}
	return false
}

func hasKSum(A []int, k, target int) bool {
	if target < 0 {
		return false
	} else if k < 0 {
		return false
	} else if len(A) > 0 && k*A[0] < target {
		return false
	} else if target == 0 && k == 0 {
		return true
	}
	// fmt.Printf("A:%v,k:%d,target:%d\n", A, k, target)
	for i := 0; i < len(A); i++ {
		if hasKSum(A[i+1:], k-1, target-A[i]) {
			return true
		}
	}
	return false
}
