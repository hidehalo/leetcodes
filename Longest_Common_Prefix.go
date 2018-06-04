package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	var min int
	for i, s := range strs {
		if i == 0 {
			min = len(s)
			continue
		}
		size := len(s)
		if size < min {
			min = size
		}
	}
	ret := ""
	for i := 0; i < min; i++ {
		var compare byte
		var ok bool
		for j, s := range strs {
			if j == 0 {
				compare = s[i]
				continue
			}
			if compare != s[i] {
				ok = false
				break
			}
			ok = true
		}
		if ok {
			ret += string(compare)
		} else {
			break
		}
	}

	return ret
}

func main() {
	fmt.Println(longestCommonPrefix([]string{
		"cbacaaa", "cbaaaa",
	}))
}
