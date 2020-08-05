package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			ret++
		}
		j++
	}
	return ret
}
