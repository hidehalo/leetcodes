package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	s := strconv.Itoa(x)
	j := len(s) - 1
	i := 0
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(2332))
}
