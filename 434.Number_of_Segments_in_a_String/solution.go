package main

func countSegments(s string) int {
	if len(s) == 1 {
		if isLetter(s[0]) {
			return 1
		}
		return 0
	}
	ret := 0
	var j int
	for i := 0; i < len(s)-1; i++ {
		if i == 0 && isLetter(s[i]) {
			ret++
			continue
		}
		j = i + 1
		if isLetter(s[i]) == false && isLetter(s[j]) == true {
			ret++
		}
	}
	return ret
}

func isLetter(b byte) bool {
	return b != ' '
}
