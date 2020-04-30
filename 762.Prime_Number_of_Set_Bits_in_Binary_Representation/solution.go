package main

import (
	"math"
)

var memo map[int]int

func countPrimeSetBits(L int, R int) int {
	memo = make(map[int]int)
	ret := 0
	for i := L; i <= R; i++ {
		if isPrimeNumber(bitsCount(i)) {
			ret++
		}
	}

	return ret
}

func bitsCount(n int) int {
	ret := 0
	for n > 0 {
		if n&0x01 == 1 {
			ret++
		}
		n >>= 1
	}

	return ret
}

func isPrimeNumber(n int) bool {
	if n < 2 {
		return false
	} else if memo[n] == 1 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	memo[n] = 1

	return true
}
