package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func optimalCoinGame(coins []int) int {
	size := len(coins)
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, size)
	}
	var x, y, z int
	for k := 0; k < size; k++ {
		for i, j := 0, k; j < size; i, j = i+1, j+1 {
			if i+2 <= j {
				// F(i+2, j)
				x = dp[i+2][j]
			}
			if i+1 <= j-1 {
				// F(i+1, j-1)
				y = dp[i+1][j-1]
			}
			if i <= j-2 {
				// F(i, j-2)
				z = dp[i][j-2]
			}
			dp[i][j] = max(coins[i]+min(x, y), coins[j]+min(y, z))
		}
	}

	return dp[0][size-1]
}

func main() {
	fmt.Println(optimalCoinGame([]int{20, 30, 2, 2, 2, 10}))
}
