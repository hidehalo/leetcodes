package main

import (
	"fmt"
)

func findComplement(num int) int {
	t := ^num
	c := 0
	r := 0
	for num > 0 {
		num = num >> 1
		r = r + t&(1<<uint(c))
		c++
	}

	return r
}

func main() {
	fmt.Println(findComplement(1))
}
