package main

func removeDuplicates(S string) string {
	return string(procedure([]byte(S)))
}

// bad
func procedure(S []byte) []byte {
	p := 0
	q := 1
	for q < len(S) {
		p = q - 1
		if S[p] == S[q] {
			t := p - 1
			k := q + 1
			for t >= 0 && k < len(S) {
				if S[t] != S[k] {
					break
				}
				t--
				k++
			}
			if t >= 0 && k < len(S) {
				S = append(S[:t+1], S[k:]...)
			} else if t < 0 && k < len(S) {
				S = S[k:]
			} else if t >= 0 && k >= len(S) {
				S = S[:t+1]
			} else {
				S = []byte{}
			}
			q = 1
		} else {
			q++
		}
	}

	return S
}
