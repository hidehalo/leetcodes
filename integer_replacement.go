package main

import (
	"fmt"
)

var dp map[int]int

func integerReplacement(n int) int {
	dp = make(map[int]int)
	for i := n; i > 1; i = next(i) {
		if v, ok := dp[i]; ok {
			dp[n] += v

			return dp[n]
		}
		dp[n]++
	}

	return dp[n]
}

func next(n int) int {
	if n%2 == 0 {
		n /= 2
	} else {
		if (n+1)%2 == 0 {
			return n + 1
		}
	}

	return n - 1
}

func main() {
	fmt.Println(integerReplacement(65535))
}
