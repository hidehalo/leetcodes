package main

import (
	"sort"
)

type Candy struct {
	key int
	num int
}

type CandySlice struct {
	data []*Candy
}

func (cs *CandySlice) Len() int {
	return len(cs.data)
}

func (cs *CandySlice) Less(i, j int) bool {
	return cs.data[i].num < cs.data[j].num
}

func (cs *CandySlice) Swap(i, j int) {
	cs.data[i], cs.data[j] = cs.data[j], cs.data[i]
}

func distributeCandies(candies []int) int {
	// greedy algo
	// sort kind key asc order by kind's count
	// plus candies of kind by order, check is sister's candies greater than brother's
	// yes: return
	// no: loop go on
	counts := make(map[int]int)
	for i := 0; i < len(candies); i++ {
		counts[candies[i]]++
	}
	candyList := make([]*Candy, 0)
	for k, v := range counts {
		candyList = append(candyList, &Candy{k, v})
	}
	sort.Sort(&CandySlice{candyList})
	kindOfSis := 0
	sumOfSis := 0
	for i := 0; i < len(candyList); i++ {
		kindOfSis++
		sumOfSis += candyList[i].num
		if sumOfSis >= len(candies)-sumOfSis {
			return kindOfSis
		}
	}

	return kindOfSis
}
