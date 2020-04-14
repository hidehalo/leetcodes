package main

func findLucky(arr []int) int {
	ret := -1
	counts := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}
	for k, v := range counts {
		if k == v {
			if v > ret {
				ret = v
			}
		}
	}
	return ret
}
