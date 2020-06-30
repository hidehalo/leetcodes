package main

var memo []int

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if memo == nil {
		memo = make([]int, 47)
		memo[1] = 1
		memo[2] = 2
		memo[3] = 3
	}
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
