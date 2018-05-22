package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	for _, word := range wordDict {
		p := strings.Index(s, word)
		if p >= 0 {
			s = s[:p] + s[len(word)+p:]
		} else {
			return false
		}
	}
	if len(s) > 0 {
		for _, str := range wordDict {
			if str == s {
				return true
			}
		}

		return false
	}

	return true
}

func main() {
	fmt.Println(wordBreak("bb", []string{"a", "b", "bbb", "bbbb"}))
}
