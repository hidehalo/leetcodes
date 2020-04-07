package main

func balancedStringSplit(s string) int {
	ret, p, q, l, r := 0, 0, 0, 0, 0
	for p < len(s) && q < len(s) {
		if s[q] == 'L' {
			l++
		} else {
			r++
		}
		q++
		if l == r {
			l, r = 0, 0
			p = q
			ret++
		}
	}

	return ret
}
