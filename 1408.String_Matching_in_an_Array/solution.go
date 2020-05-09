package main

import "strings"

func stringMatching(words []string) []string {
	ret := make([]string, 0)
	for i := range words {
		for j := range words {
			if i == j {
				continue
			}
			if strings.Index(words[j], words[i]) != -1 {
				ret = append(ret, words[i])
				break
			}
		}
	}

	return ret
}
