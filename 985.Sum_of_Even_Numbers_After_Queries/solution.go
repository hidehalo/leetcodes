package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	ret := make([]int, 0)
	for _, q := range queries {
		v := q[0]
		i := q[1]
		A[i] += v
		evenSum := 0
		for k := 0; k < len(A); k++ {
			if A[k]%2 == 0 {
				evenSum += A[k]
			}
		}
		ret = append(ret, evenSum)
	}
	return ret
}
