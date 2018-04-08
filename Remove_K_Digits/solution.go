package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	len := len(num)
	if k >= len {
		return "0"
	}
	for k > 0 {
		for i := 0; i < len-1; i++ {
			if num[i] > num[i+1] {
				num = strings.Replace(num, string(num[i]), "", 1)
				len--
				k--
				if string(num[0]) == "0" && len > 1 {
					num = strings.Replace(num, string(num[0]), "", 1)
					len--
				}
				break
			} else if i == len-2 {
				num = strings.Replace(num, string(num[i+1]), "", 1)
				len--
				k--
				break
			}
		}
	}

	return num
}

func main() {
	fmt.Println(removeKdigits("10200", 1))
}
