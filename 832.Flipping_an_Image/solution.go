package main

func flipAndInvertImage(A [][]int) [][]int {
	m := len(A)
	n := len(A[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			//reverse: 0 <-> 1 = 1 0
			// invert: 1 0 = 0 1
			if A[i][j] != A[i][n-j-1] {
				continue
			}
			A[i][j], A[i][n-j-1] = 1^A[i][j], 1^A[i][j]
		}
		if n%2 != 0 {
			A[i][n/2] = 1 ^ A[i][n/2]
		}
	}

	return A
}
