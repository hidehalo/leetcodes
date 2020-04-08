package main

import (
	"sort"
)

func sortString(s string) string {
	// convert to int slice O(n)
	ui8s := make([]int, 0)
	for j := 0; j < len(s); j++ {
		ui8s = append(ui8s, int(s[j]))
	}
	// sort order ASC O(nlgn)
	sort.Ints(ui8s)
	counts := make(map[int]int)
	// note: double link list is better than hash table
	ranks := make(map[int]int)
	current := -1
	rank := 0 // init rank
	// build ranks&counts hash table O(n)
	for i := 0; i < len(ui8s); i++ {
		if current != ui8s[i] {
			current = ui8s[i]
			rank++
			ranks[rank] = current
		}
		counts[current]++
	}
	// note: trick for unique character string
	if len(ranks) == 1 {
		return s
	}
	ret := make([]byte, 0)
	p := 1     // rank position
	order := 0 // 0: ASC|1: DESC
	// O(n)
	for i := 0; i < len(ui8s); i++ {

		if p == len(ranks)+1 {
			order = 1
		} else if p == 0 {
			order = 0
		}

		if counts[ranks[p]] <= 0 {
			if p == len(ranks)+1 {
				order = 1
			} else if p == 0 {
				order = 0
			}
		}

		counts[ranks[p]]--
		ret = append(ret, byte(ranks[p]))

		if order == 1 {
			p--
		} else {
			p++
		}
	}

	return string(ret)
}
