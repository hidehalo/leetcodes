package main

func findPaths(m int, n int, N int, i int, j int) int {
	dp := make([]int, m*n)
	for k := 0; k < m; k++ {
		dp[k] = 1
		dp[(n-1)*m+k] = 1
	}
	for k := 0; k < n; k++ {
		dp[k] = 1
		dp[(m-1)*n+k] = 1
	}

	// return sum of d([i,j], [boundary points]) <= N
}
