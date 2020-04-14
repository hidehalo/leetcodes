package main

func reverseWords(s string) string {
	j := 0
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' || i == len(b)-1 {
			m := j
			n := i
			for m < n {
				if b[n] == ' ' {
					n--
					continue
				}
				b[m], b[n] = b[n], b[m]
				m++
				n--
			}
			j = i + 1
		}
	}
	return string(b)
}
