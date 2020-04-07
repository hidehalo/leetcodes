package main

func maximum69Number(num int) int {
	base := 10000
	p := num
	for base > 0 {
		if p > base && p/base < 9 {
			return num + 3*base
		}
		p %= base
		base /= 10
	}
	return num
}
