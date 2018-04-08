package main

import (
	"fmt"
)

func uniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	cnt := make(map[string]int)
	for _, word := range words {
		tmp := ""
		for _, alpha := range word {
			tmp += morseCodes[alpha-97]
		}
		if _, ok := cnt[tmp]; !ok {
			cnt[tmp] = 1
		}
	}

	return len(cnt)
}

func main() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
