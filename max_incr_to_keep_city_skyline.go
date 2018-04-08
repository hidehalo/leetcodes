package main

import (
	"fmt"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	max := make(map[string][]int)
	for i, vi := range grid {
		if _, ok := max["height"]; !ok {
			max["height"] = make([]int, len(vi))
			max["width"] = make([]int, len(grid))
		}
		for j, vj := range vi {
			if vj > max["height"][j] {
				max["height"][j] = vj
			}
			if vj > max["width"][i] {
				max["width"][i] = vj
			}
		}
	}
	ret := 0
	for i, vi := range grid {
		for j, vj := range vi {
			if vj < max["width"][i] && vj < max["height"][j] {
				diff := 0
				if max["width"][i] < max["height"][j] {
					diff = (max["width"][i] - vj)
				} else {
					diff = (max["height"][j] - vj)
				}
				ret += diff
			}
		}
	}

	return ret
}

func main() {
	arr := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(arr))
}
