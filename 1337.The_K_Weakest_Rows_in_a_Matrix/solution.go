package main

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	counts := make([]int, len(mat))
	vmap := make(map[int][]int)

	for i := 0; i < len(mat); i++ {
		s := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				break
			}
			s++
		}
		counts[i] = s
		vmap[s] = append(vmap[s], i)
	}

	sort.Ints(counts)
	ret := make([]int, 0)
	i := 0
	for k > 0 {
		for j := 0; j < len(vmap[i]); j++ {
			ret = append(ret, vmap[i][j])
			k--
			if k <= 0 {
				break
			}
		}
		i++
	}

	return ret
}
