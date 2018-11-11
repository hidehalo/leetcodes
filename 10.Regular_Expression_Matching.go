package main

import "fmt"

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	if p == "*" {
		return false
	}
	sizeS, sizeP := len(s), len(p)
	if sizeS <= 0 {
		return false
	}
	if sizeP <= 0 {
		return false
	}
	i, j := 0, 0
	var cache byte
	for i < sizeS && j < sizeP {
		if p[j] == '.' {
			if j < sizeP-1 && p[j+1] == '*' {
				return true
			}
			j++
			i++
			continue
		} else if j < sizeP-1 && p[j+1] == '*' {
			// for i < sizeS && s[i] == p[j] {
			// 	i++
			// }
			// j += 2
			// if i >= sizeS && j < sizeP {
			// 	fmt.Println(sizeP-j, i)
			// 	if sizeP-j > i {
			// 		return false
			// 	}
			// 	for j < sizeP && s[sizeS-1] == p[j] {
			// 		j++
			// 	}
			// }
			for i < sizeS && s[i] == cache {
				i++
			}
			cache = p[j]
			if s[i] == p[j] {
				i++
			}
			j += 2
			continue
		}
		if s[i] != p[j] {
			if s[i] == cache {
				i++
			} else {
				return false
			}
		} else {
			i++
			j++
		}
	}

	return true
}

func main() {
	fmt.Println(isMatch("aaa", "a*aaaaaaa"))
}
