package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	var i, j, ret int
	size := len(haystack)
	sizeP := len(needle)
	for i = 0; i < size; i++ {
		if j >= sizeP {
			break
		} else if j > 0 && haystack[i] != needle[j] {
			j = 0
			i = ret
		}
		if haystack[i] == needle[0] {
			ret = i
		}
		if haystack[i] == needle[j] {
			j++
		}
	}
	if j == sizeP {
		return ret
	}

	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}
