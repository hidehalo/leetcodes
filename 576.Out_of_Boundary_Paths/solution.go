package main

func findPaths(m int, n int, N int, i int, j int) int {
	y := n - j
	if y < 0 {
		y = -y
	}
	x := m - i
	if x < 0 {
		x = -x
	}
	b := make([]int, m)
	for k:=0;k<m;k++{
		b[k] := make([]int,n)
	}
	// Set shortest path distance for boundary points
	for k := 0; k < m; k++ {
		b[0][k] = j+1
		b[k][0]=y+k
	}
	for k := 0; k < n; k++ {
		b[k][n] = i + 1
		b[n][k] = x + k
	}

	// return sum of d([i,j], [boundary points]) <= N
}
