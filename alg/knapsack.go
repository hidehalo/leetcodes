package main

import (
	"fmt"
)

var cnt, cntDP int

func knapsack(W int, w []int, v []int, n int) int {
	cnt++
	if W <= 0 || n <= 0 {
		return 0
	}
	if W < w[n-1] {
		return knapsack(W, w, v, n-1)
	}

	a := v[n-1] + knapsack(W-w[n-1], w, v, n-1)
	b := knapsack(W, w, v, n-1)
	if a > b {
		return a
	}

	return b
}

func knapsackDP(W int, w []int, v []int, n int) int {
	var B [][]int
	B = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		B[i] = make([]int, W+1)
	}
	for i := 1; i <= n; i++ {
		for wi := 0; wi <= W; wi++ {
			cntDP++
			if w[i-1] <= W {
				if v[i-1]+B[i-1][W-w[i-1]] > B[i-1][wi] {
					B[i][wi] = v[i-1] + B[i-1][W-w[i-1]]
				} else {
					B[i][wi] = B[i-1][wi]
				}
			} else {
				B[i][wi] = B[i-1][wi]
			}
		}
	}

	return B[n][W]
}

func main() {
	w := []int{10, 20, 30}
	b := []int{60, 100, 120}
	n := len(w)
	fmt.Println("non-dp Result:", knapsack(50, w, b, n))
	fmt.Println("non-dp T(n):", cnt)
	fmt.Println("dp Result:", knapsackDP(50, w, b, n))
	fmt.Println("dp T(n):", cntDP)
}
