package main

func minDeletionSize(A []string) int {
	matrix := make([][]byte, 0)
	for i := 0; i < len(A); i++ {
		matrix = append(matrix, []byte(A[i]))
	}

	ret := 0
	for j := 0; j < len(matrix[0]); j++ {
		desc := 0
		for i := 0; i < len(matrix)-1; i++ {
			// the real question is count decreasing order columns
			// what is decreasing order column? any of c[i][j] > c[i+1][j] i:[0,m-2],j:[0,n-1]
			if matrix[i][j] > matrix[i+1][j] {
				desc = 1
				break
			}
		}
		if desc == 1 {
			ret++
		}
	}

	return ret
}
