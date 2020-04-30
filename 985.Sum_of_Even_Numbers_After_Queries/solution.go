package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	ret := make([]int, 0)
	evenSum := 0
	for k := 0; k < len(A); k++ {
		if A[k]%2 == 0 {
			evenSum += A[k]
		}
	}
	for _, q := range queries {
		v := q[0]
		i := q[1]
		if A[i]%2 == 0 && (A[i]+v)%2 == 0 {
			evenSum += v
		} else if A[i]%2 != 0 && (A[i]+v)%2 == 0 {
			evenSum += (A[i] + v)
		} else if A[i]%2 == 0 && (A[i]+v)%2 != 0 {
			evenSum -= A[i]
		}
		A[i] += v
		ret = append(ret, evenSum)
	}
	return ret
}
