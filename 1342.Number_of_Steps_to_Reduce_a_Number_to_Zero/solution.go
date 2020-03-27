package main

var dp map[int]int

func numberOfSteps(num int) int {
	if dp == nil {
		dp = make(map[int]int)
	}

	if v, ok := dp[num]; ok {
		return v
	}

	p := num
	for p > 0 {
		if p%2 == 0 {
			p >>= 1
		} else {
			p--
		}
		if v, ok := dp[p]; ok {
			dp[num] += v + 1
			break
		}
		dp[num]++
	}

	return dp[num]
}
