package main

var memo map[int]int

func countPrimes(n int) int {
	if memo == nil {
		memo = make(map[int]int)
		memo[2] = 1
	}
	delta := 0
	p := n
	for {
		if _, ok := memo[p]; ok == true || p <= 2 {
			break
		}
		if isPrime(p) == true {
			delta++
		}
		p--
	}
	memo[n] = memo[p] + delta

	return memo[n]
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
