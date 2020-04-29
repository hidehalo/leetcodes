package main

func toGoatLatin(S string) string {
	vowels := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	ret := make([]byte, 0)
	checkFirst := true
	endMa := false
	wordCount := 0
	spec := false
	var moveToEnd byte
	for i := 0; i < len(S); i++ {
		if checkFirst == true {
			checkFirst = false
			endMa = true
			if !(vowels[S[i]] > 0 ||
				(i < len(S)-1 &&
					(S[i] == 'o' || S[i] == 'O') &&
					(S[i+1] == 'r' || S[i+1] == 'R'))) {
				moveToEnd = S[i]
				if len(S)-1 != i {
					continue
				} else {
					spec = true
				}
			}
		}
		if len(S)-1 == i && spec == false {
			ret = append(ret, S[i])
		}
		if S[i] == ' ' || len(S)-1 == i {
			if moveToEnd != 0 {
				ret = append(ret, moveToEnd)
				moveToEnd = 0
			}
			wordCount++
			checkFirst = true
			if endMa == true {
				endMa = false
				ret = append(ret, 'm', 'a')
			}
			for i := 0; i < wordCount; i++ {
				ret = append(ret, 'a')
			}
		}
		if len(S)-1 != i {
			ret = append(ret, S[i])
		}
	}

	return string(ret)
}
