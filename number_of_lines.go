package main

import (
	"fmt"
)

func numberOfLines(widths []int, S string) []int {
	if len(S) <= 0 {
		return []int{0, 0}
	}
	units := 0
	lines := 1
	for _, alpha := range S {
		step := widths[alpha-97]
		units += step
		flag := units / 100
		next := units % 100
		if flag > 0 && next > 0 {
			lines++
			units = step
		} else if next == 0 {
			lines++
			units = 0
		}
	}

	return []int{lines, units}
}

func main() {
	w := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "bbbcccdddaaa"
	ret := numberOfLines(w, s)

	fmt.Println(ret)
}
