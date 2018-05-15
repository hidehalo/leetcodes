package main

import (
	"fmt"
)

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0, 10000)
	for i := left; i <= right; i++ {
		if helper(i) == true {
			ret = append(ret, i)
		}
	}

	return ret
}

func helper(n int) bool {
	t := n
	for t > 0 {
		d := t % 10
		if d == 0 || n%d != 0 {
			return false
		}
		t /= 10
	}

	return true
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
