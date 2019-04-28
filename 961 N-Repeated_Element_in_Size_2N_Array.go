package main

import "fmt"

func repeatedNTimes(A []int) int {
	check := make(map[int]int)
	for _, a := range A {
		if c, _ := check[a]; true {
			if c >= 1 {
				return a
			}
			check[a]++
		}
	}
	return -1
}

func main() {
	num := []int{1, 2, 3, 2, 5, 2}
	fmt.Println(repeatedNTimes(num))
}
