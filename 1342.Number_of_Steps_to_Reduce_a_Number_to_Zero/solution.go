package main

var dp map[int]int

func numberOfSteps(num int) int {
	if dp == nil {
		dp = make(map[int]int)
		dp[1] = 1
	}
	q, p := 1, 1
	if v, ok := dp[num]; ok {
		return v
	}
	for p < num {
		q = p
		if p%2 == 0 { //fixme: wrong algo
			p++
		} else {
			p *= 2
		}
		dp[p] = dp[q] + 1
	}
	if p == num {
		return dp[p]
	}
	return dp[q]
}
