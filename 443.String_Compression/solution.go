package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	ret := 0
	var currentByte byte
	var continuous int
	var write int
	for i := 0; i < len(chars); i++ {
		if continuous == 0 {
			currentByte = chars[i]
			continuous++
			ret++
			continue
		}
		if currentByte == chars[i] {
			continuous++
		} else {
			fmt.Printf("字符%c连续%d次后被终止\n", currentByte, continuous)
			chars[write] = currentByte
			write++
			if continuous > 1 {
				continuousStr := strconv.Itoa(continuous)
				for j := 0; j < len(continuousStr); j++ {
					chars[write] = continuousStr[j]
					write++
				}
			}
			ret += itoaCount(continuous) + 1
			currentByte = chars[i]
			continuous = 1
		}
	}
	chars[write] = currentByte
	write++
	fmt.Printf("字符%c连续%d次后被终止\n", currentByte, continuous)
	if continuous > 1 {
		continuousStr := strconv.Itoa(continuous)
		for j := 0; j < len(continuousStr); j++ {
			chars[write] = continuousStr[j]
			write++
		}
	}
	chars = chars[:ret+itoaCount(continuous)]
	fmt.Println(string(chars))
	return ret + itoaCount(continuous)
}

func itoaCount(n int) int {
	if n <= 1 {
		return 0
	}
	ret := 0
	for n > 0 {
		ret++
		n /= 10
	}
	return ret
}
