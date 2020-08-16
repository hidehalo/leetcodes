package main

func longestPalindrome(s string) int {
	bucket := make([][]int, 2)
	for i := 0; i < 2; i++ {
		bucket[i] = make([]int, 27)
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			bucket[0][int(s[i]-'a')]++
		} else {
			bucket[1][int(s[i]-'A')]++
		}
	}
	ret := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 27; j++ {
			if bucket[i][j] > 1 {
				if bucket[i][j]%2 == 0 {
					ret += bucket[i][j]
				} else {
					ret += bucket[i][j] - 1
				}
			}
		}
	}
	if ret < len(s) {
		return ret + 1
	}
	return ret
}
