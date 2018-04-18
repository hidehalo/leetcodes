package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	words := strings.Split(paragraph, " ")
	cnt := make(map[string]int)
	for _, peach := range words {
		in := false
		peach = strings.ToLower(strings.Trim(peach, ",.!?;'\""))
		for _, ban := range banned {
			if peach == ban {
				in = true
				break
			}
		}
		if !in {
			cnt[peach]++
		}
	}
	max := 0
	ret := ""
	for k, v := range cnt {
		if max < v {
			max = v
			ret = k
		}
	}

	return ret
}

func main() {
	banned := []string{"hit"}
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", banned))
}
