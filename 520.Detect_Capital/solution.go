package main

func detectCapitalUse(word string) bool {
	var mode int
	for i := 0; i < len(word); i++ {
		if mode == 0 {
			if word[i] >= 'A' && word[i] <= 'Z' {
				mode = 3
			} else if word[i] >= 'a' && word[i] <= 'z' {
				mode = 2
			}
			if mode == 3 {
				if i == len(word)-1 || word[i+1] >= 'A' && word[i+1] <= 'Z' {
					mode = 1
				}
			}
			continue
		} else if word[i] == ' ' {
			mode = 0
			continue
		}

		if (mode == 3 || mode == 2) && !(word[i] >= 'a' && word[i] <= 'z') {
			return false
		}
		if mode == 1 && !(word[i] >= 'A' && word[i] <= 'Z') {
			return false
		}
	}
	return true
}
