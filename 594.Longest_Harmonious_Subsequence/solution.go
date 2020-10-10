package main

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	lhs := 0
	for k, v := range m {
		if m[k-1] == 0 {
			continue
		}
		if v+m[k-1] > lhs {
			lhs = v + m[k-1]
		}
	}
	return lhs
}
