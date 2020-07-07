package main

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	diff := make([][]byte, 0)
	alphabets := make([]int, 26)
	for i := 0; i < len(A); i++ {
		alphabets[int(A[i]-'a')]++
		if A[i] != B[i] {
			diff = append(diff, []byte{A[i], B[i]})
		}
		if len(diff) > 2 {
			return false
		}
	}
	if len(diff) == 1 {
		return false
	} else if len(diff) == 0 {
		for j := 0; j < len(alphabets); j++ {
			if alphabets[j] > 1 {
				return true
			}
		}
		return false
	}

	return diff[0][0] == diff[1][1] && diff[0][1] == diff[1][0]
}
