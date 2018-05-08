package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	f1 := 0
	f2 := 0
	f := 0
	size := len(cost)
	for i := size - 1; i >= 0; i-- {
		var min int
		if f1 < f2 {
			min = f1
		} else {
			min = f2
		}
		f = cost[i] + min
		f2 = f1
		f1 = f
	}

	if f1 < f2 {
		return f1
	}

	return f2
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
