package main

func trailingZeroes(n int) int {
	ret := 0
	for n >= 5 {
		ret += n / 5
		n /= 5
	}

	return ret
}
