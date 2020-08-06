package main

func findTheDifference(s string, t string) byte {
	ha := make([]int, 26)
	ht := make([]int, 26)
	for i := 0; i < len(t); i++ {
		ht[int(t[i]-'a')]++
		if i < len(s) {
			ha[int(s[i]-'a')]++
		}
	}
	for i := 0; i < 26; i++ {
		if ha[i] != ht[i] {
			return 'a' + byte(i)
		}
	}
	return ' '
}
