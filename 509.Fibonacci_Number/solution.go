package main

var memo []int

func fib(N int) int {
	if memo == nil {
		memo = make([]int, 0)
		memo = append(memo, 0, 1)
	}
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	} else if N < len(memo)-1 {
		return memo[N-1] + memo[N-2]
	}
	for i := len(memo); i <= N; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}

	return memo[N-1] + memo[N-2]
}
