package main

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
