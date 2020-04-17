package main

var memo []int

func fib(N int) int {
	if memo == nil {
		memo = make([]int, 0)
		memo = append(memo, 0, 1)
	}
	if N < len(memo)-1 {
		return memo[N]
	}
	for i := 2; i <= N; i++ {
		val = memo[i-1] + memo[i-2]
		memo = append(memo, val)
	}

	return memo[N]
}
