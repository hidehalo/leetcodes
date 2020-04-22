package main

import "fmt"

func removeDuplicates(S string) string {
	p := 0
	q := 1
	for ; q < len(S); q++ {
		p = q - 1
		if S[p] == S[q] {
			t := p
			k := q
			for t >= 0 && k < len(S) {
				if S[t] != S[k] {
					break
				}
				t--
				k++
			}
			if t >= 0 {
				S = S[t:p] + S[k:]
			} else {
				S = S[k:]
			}
			fmt.Println(S)
		}
	}

	return S
}
