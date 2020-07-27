package main

func titleToNumber(s string) int {
	ret := 0
	weight := 1
	zeroVal := int('@')
	for i := len(s) - 1; i >= 0; i-- {
		ret += weight * (int(s[i]) - zeroVal)
		weight *= 26
	}
	return ret
}
