package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := prices[0]
	profit := 0
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		buy = min(buy, prices[i])
		profit = max(profit, prices[i]-buy)
		if prices[i] < prices[i-1] {
			maxProfit += profit
			profit = 0
			buy = prices[i]
		}
	}
	return maxProfit + profit
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
