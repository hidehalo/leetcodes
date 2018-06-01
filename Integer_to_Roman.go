package main

import (
	"fmt"
)

func intToRoman(num int) string {
	table := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	ret := ""
	i := 0
	for num > 0 {
		v := keys[i]
		if num >= v {
			num -= v
			ret += table[v]
		} else {
			i++
		}
	}

	return ret
}

func main() {
	fmt.Println(intToRoman(1994))
}
