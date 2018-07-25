package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i, c := range haystack {
		if i+len(needle) > len(haystack) {
			break
		} else if byte(c) == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
	// hSize := len(haystack)
	// nSize := len(needle)
	// offsetIndecies := make([]int, 0, hSize)
	// for i, c := range haystack {
	// 	if byte(c) == needle[0] && i+nSize <= hSize {
	// 		offsetIndecies = append(offsetIndecies, i)
	// 	}
	// }
	// ret := -1
	// if len(offsetIndecies) > 0 {
	// 	for _, offset := range offsetIndecies {
	// 		ret = conquer(haystack, offset, needle)
	// 		if ret >= 0 {
	// 			break
	// 		}
	// 	}
	// }

	// return ret
}

// func conquer(haystack string, offset int, needle string) int {
// 	for i, c := range needle {
// 		if haystack[offset+i] != byte(c) {
// 			return -1
// 		}
// 	}

// 	return offset
// }

func main() {
	fmt.Println(strStr("hello", "ll"))
}
