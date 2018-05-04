package main

import (
	"fmt"
)

func findAnagrams(s string, p string) []int {
	sizeP := len(p)
	sizeS := len(s)
	if sizeS < sizeP {
		return []int{}
	}
	ret := make([]int, 0, sizeS-1)
	hashMapP := make(map[byte]int)
	hashMapS := make(map[byte]int)
	for i := range s {
		if i < sizeP {
			hashMapP[p[i]]++
			hashMapS[s[i]]++
		} else {
			hashMapS[s[i]]++
			hashMapS[s[i-sizeP]]--
		}
		ok := true
		for k, v := range hashMapP {
			if hashMapS[k] != v {
				ok = false
				break
			}
		}
		tmp := i - sizeP + 1
		if ok && tmp >= 0 {
			ret = append(ret, tmp)
		}
	}

	return ret
}

func main() {
	fmt.Println(findAnagrams("abab", "ab"))
}
