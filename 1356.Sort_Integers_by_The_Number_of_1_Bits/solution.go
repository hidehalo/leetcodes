package main

import (
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Ints(arr)
	buckets := make(map[int][]int)
	maxBitsNum := 128
	for i := 0; i < len(arr); i++ {
		bitsCount := countBits(arr[i])
		if bitsCount > maxBitsNum {
			maxBitsNum = bitsCount
		}
		buckets[bitsCount] = append(buckets[bitsCount], arr[i])
	}

	ret := make([]int, 0)
	for i := 0; i <= maxBitsNum; i++ {
		if _, ok := buckets[i]; !ok {
			continue
		}
		for j := 0; j < len(buckets[i]); j++ {
			ret = append(ret, buckets[i][j])
		}
	}

	return ret
}

func countBits(n int) int {
	ret := 0
	for n > 0 {
		if n&0x01 == 1 {
			ret++
		}
		n >>= 1
	}
	return ret
}
