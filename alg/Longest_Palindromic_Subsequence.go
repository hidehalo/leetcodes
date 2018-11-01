package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func longestPalindromicSubsequence(str string) int {
	size := len(str)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = append(make([]int, size))
		dp[i][i] = 1
	}
	for subseqSize := 2; subseqSize <= size; subseqSize++ {
		for i := 0; i < size-subseqSize+1; i++ {
			j := i + subseqSize - 1
			if str[i] == str[j] && subseqSize == 2 {
				dp[i][j] = 2
			} else if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][size-1]
}

func main() {
	fmt.Println(longestPalindromicSubsequence("whatisthefuckkcufehtsitahwlalala~"))
}
