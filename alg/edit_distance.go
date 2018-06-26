package main

import (
	"fmt"
)

var cnt, cntDP int

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func editDistance(s []rune, t []rune, m int, n int) int {
	cnt++
	if m <= 0 {
		return n
	}
	if n <= 0 {
		return m
	}
	if s[m-1] == t[n-1] {
		return editDistance(s, t, m-1, n-1)
	}

	insert := editDistance(s, t, m, n-1)
	remove := editDistance(s, t, m-1, n)
	replace := editDistance(s, t, m-1, n-1)

	min := 0
	if insert > remove {
		min = remove
	} else {
		min = insert
	}
	if replace < min {
		min = replace
	}

	return 1 + min
}

func editDistanceDP(s []rune, t []rune, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			cntDP++
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	s := []rune("hello")
	t := []rune("world")
	ss := len(s)
	st := len(t)
	fmt.Println("non-dp Result:", editDistance(s, t, ss, st))
	fmt.Println("dp Result:", editDistanceDP(s, t, ss, st))
	fmt.Println("non-dp T(n):", cnt)
	fmt.Println("dp T(n):", cntDP)
}
