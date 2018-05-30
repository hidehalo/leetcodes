package main

import (
	"fmt"
)

func romanToInt(s string) int {
	t := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	size := len(s)
	ret := 0
	for i := 0; i < size; i++ {
		if i+1 < size {
			if t[s[i+1]] > t[s[i]] {
				ret -= t[s[i]]
			} else {
				ret += t[s[i]]
			}
		} else {
			ret += t[s[i]]
		}
	}

	return ret
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
