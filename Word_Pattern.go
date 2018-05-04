package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")
	hashMap := make(map[byte]string)
	dict := make(map[string]byte)
	if len(words) != len(pattern) {
		return false
	}
	for i, word := range words {
		if _, ok := hashMap[pattern[i]]; !ok {
			hashMap[pattern[i]] = word
			if _, ok := dict[word]; !ok {
				dict[word] = pattern[i]
			} else if dict[word] != pattern[i] {
				return false
			}
		} else if hashMap[pattern[i]] != word {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
