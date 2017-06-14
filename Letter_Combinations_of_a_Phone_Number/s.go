package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	button := map[string]string{
		"1": "",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
		"0": "",
		"#": "",
	}
	ret := make([]string, 1)
	for _, b := range digits {
		chars := button[string(b)]
		charLen := len(chars)
		if charLen <= 0 {
			continue
		}
		tmp := make([]string, 0)
		for i := 0; i < charLen; i++ {
			for j := 0; j < len(ret); j++ {
				tmp = append(tmp, ret[j]+string(chars[i]))
			}
		}
		ret = tmp
	}

	return ret
}

func main() {
	str := "23"
	ret := letterCombinations(str)
	fmt.Println(ret)
}
