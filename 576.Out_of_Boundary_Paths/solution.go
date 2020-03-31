package main

var dp map[int]int

func hash(m, n, N, i, j int) int {
	return m<<24 | n<<18 | N<<12 | i<<6 | j
	// return m*1e8 + n*1e6 + N*1e4 + i*1e2 + j
}

func findPaths(m int, n int, N int, i int, j int) int {
	if dp == nil {
		dp = make(map[int]int)
	}
	if v, ok := dp[hash(m, n, N, i, j)]; ok { // check memo
		return v
	} else if ((i == m && j < n && j >= 0) ||
		(j == n && i < m && i >= 0) ||
		(i == -1 && j < n && j >= 0) ||
		(j == -1 && i < m && i >= 0)) && N >= 0 { // bingo

		return 1
	} else if N == 0 { // overstep N but not out of boundary
		return 0
	}

	dp[hash(m, n, N, i, j)] = (findPaths(m, n, N-1, i-1, j) + // move down
		// move up
		findPaths(m, n, N-1, i+1, j) +
		// move left
		findPaths(m, n, N-1, i, j-1) +
		// move right
		findPaths(m, n, N-1, i, j+1)) % (1e9 + 7)

	return dp[hash(m, n, N, i, j)]
}
