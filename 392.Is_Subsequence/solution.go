package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	} else if len(s) == 0 {
		return true
	}
	var p, q int
	for p < len(s) {
		for q < len(t) && t[q] != s[p] {
			q++
		}
		if q == len(t) {
			return false
		} else if t[q] == s[p] {
			q++
		}
		p++
	}

	return true
}
