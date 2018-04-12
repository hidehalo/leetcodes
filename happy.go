package main

import (
	"fmt"
)

var memo = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
var sums = make(map[int]int)
var dp = make(map[int]bool)

func isHappy(n int) bool {
	if v, ok := dp[n]; ok {
		return v
	}
	loop := 0
	for i := subProgram(n); i != 1; i = subProgram(i) {
		loop++
		if loop > 20 {
			dp[n] = false

			return false
		}
	}
	dp[n] = true

	return true
}

func subProgram(n int) int {
	sum := 0
	if v, ok := sums[n]; ok {
		return v
	}
	for mode := n; mode > 0; {
		tmp := mode % 10
		sum += memo[tmp]
		mode /= 10
	}
	sums[n] = sum

	return sum
}

func main() {
	num := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		num = append(num, i)
	}
	for _, v := range num {
		fmt.Println(v, ":", isHappy(v))
	}
	// fmt.Println(7, ":", isHappy(7))
}
