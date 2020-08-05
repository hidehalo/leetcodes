package main

func repeatedSubstringPattern(s string) bool {
	for divide := 2; divide <= len(s); divide++ {
		// Choose a size that can be divided equally
		if len(s)%divide != 0 {
			continue
		}
		step := len(s) / divide
		offset := step
		// Test each substring
		var i int
		for i = offset; i < len(s); i += step {
			if strEq(s[0:offset], s[i:i+step]) != true {
				break
			}
		}
		if i == len(s) {
			return true
		}
	}

	return false
}

func strEq(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
