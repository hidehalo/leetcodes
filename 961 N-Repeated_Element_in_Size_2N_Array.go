package main

import "fmt"

func repeatedNTimes(A []int) int {
	var a, b int
	for _, num := range A {
		b = num
		fmt.Println(a, b)
		if a|b-a^b == b {
			return b
		}
		a = a | b
	}

	return -1
}

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	fmt.Println(repeatedNTimes(num))
}
