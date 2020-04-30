package main

func isToeplitzMatrix(matrix [][]int) bool {
	M := len(matrix)
	if M == 0 {
		return false
	}
	N := len(matrix[0])
	for i := 0; i < M; i++ {
		if checkToeplitz(matrix, M, N, i, 0) == false {
			return false
		}
	}
	for j := 0; j < N; j++ {
		if checkToeplitz(matrix, M, N, 0, j) == false {
			return false
		}
	}

	return true
}

func checkToeplitz(matrix [][]int, M, N, offsetM, offsetN int) bool {
	// fmt.Println("Start testing.....")
	i, j := offsetM, offsetN
	checkValue := matrix[offsetM][offsetN]
	for i < M && j < N {
		// fmt.Printf("Check matrix[%d][%d] == matrix[%d][%d]? %v\n", i+1, j+1, offsetM+1, offsetN+1, matrix[i][j] == checkValue)
		if matrix[i][j] != checkValue {
			return false
		}
		i++
		j++
	}

	return true
}
