package main

func reverseVowels(s string) string {
	prev, next := 0, 0
	data := []byte(s)
	for prev < len(data) {
		if isVowels(data[prev]) {
			for next = prev + 1; next < len(data); next++ {
				if isVowels(data[next]) {
					data[prev], data[next] = data[next], data[prev]
					break
				}
			}
			prev = next
		} else {
			prev++
		}
	}

	return string(data)
}

func isVowels(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
		b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
