package main

func reverseVowels(s string) string {
	prev, next := 0, len(s)-1
	data := []byte(s)
	for prev < next {
		firstCheck := isVowels(data[prev])
		secondCheck := isVowels(data[next])
		if firstCheck && secondCheck {
			data[prev], data[next] = data[next], data[prev]
			next--
			prev++
		} else if firstCheck {
			next--
		} else if secondCheck {
			prev++
		} else {
			next--
			prev++
		}
	}

	return string(data)
}

func isVowels(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
