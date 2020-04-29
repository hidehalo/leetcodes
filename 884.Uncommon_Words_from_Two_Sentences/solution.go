package main

import "strings"

func uncommonFromSentences(A string, B string) []string {
	uniq := make(map[string]int)
	C := A + " " + B
	words := strings.Split(C, " ")
	for _, word := range words {
		uniq[word]++
	}
	ret := make([]string, 0)
	for k, v := range uniq {
		if v == 1 {
			ret = append(ret, k)
		}
	}

	return ret
}
