package main

import (
	"fmt"
)

func reverse(x int) int {
	flag := -1
	if x >= 0 {
		flag = 1
	}
	t := x * flag
	x = 0
	for t > 0 {
		x = x*10 + t%10
		t /= 10
	}
	r := x * flag
	if r > 1<<31-1 || r < -1<<31 {
		return 0
	}

	return r
}

func main() {
	fmt.Println(reverse(1534236469))
}
