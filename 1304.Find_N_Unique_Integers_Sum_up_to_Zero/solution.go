package main

func sumZero(n int) []int {
	ret := make([]int, 0)
	for i := 1; i < n; i += 2 {
		ret = append(ret, i, -i)
	}
	if n%2 != 0 {
		ret = append(ret, 0)
	}
	return ret
}
