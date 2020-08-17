package main

func isPalindrome(s string) bool {
	p, q := 0, len(s)-1
	pLen := 0
	for p <= q {
		if !isLetter(s[p]) && !isNumber(s[p]) {
			p++
			continue
		}
		if !isLetter(s[q]) && !isNumber(s[q]) {
			q--
			continue
		}
		if toLower(s[p]) != toLower(s[q]) {
			return false
		}

		if p != q {
			pLen++
		} else {
			pLen += 2
		}

		p++
		q--
	}

	return pLen >= 0
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
