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
	if sizeP == sizeS+1 &&
		p[sizeP-1] != '*' {
		return false

	}
	i, j := 0, 0
	for i < sizeS && j < sizeP {
		if p[j] == '.' {
			if j < sizeP-1 && p[j+1] == '*' {
				return true
			}
			j++
			i++
			continue
		} else if j < sizeP-1 && p[j+1] == '*' {
			for i < sizeS && s[i] == p[j] {
				i++
			}
			j += 2
			continue
		}
		i++
		j++
	}
	fmt.Println(i, j)
	if i < sizeS || j < sizeP {
		return false
	}
	return true
}

func main() {
	fmt.Println(isMatch("aaa", "a*aa"))
}
