package main

func checkRecord(s string) bool {
	count := make([]int, 2)
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			if count[0] >= 1 {
				return false
			}
			count[0]++
		} else if s[i] == 'L' {
			if i != 0 && s[i-1] != 'L' {
				count[1] = 0
			}
			if count[1] >= 2 {
				return false
			}
			count[1]++
		}
	}
	return true
}
