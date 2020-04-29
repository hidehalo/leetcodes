package main

func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}
	ret := make([][]int, len(A[0]))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if len(ret[j]) == 0 {
				ret[j] = make([]int, len(A))
			}
			ret[j][i] = A[i][j]
		}
	}
	return ret
}
