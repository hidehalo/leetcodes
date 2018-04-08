package main

import (
	"fmt"
)

func largestTriangleArea(points [][]int) float64 {
	max := []int{0, 0}
	min := []int{2e10, 2e10}
	for _, p := range points {
		x := p[0]
		y := p[1]
		if x > max[0] {
			max[0] = x
		}
		if x < min[0] {
			min[0] = x
		}
		if y > max[1] {
			max[1] = y
		}
		if y < min[1] {
			min[1] = y
		}
	}

	return float64((max[0]-min[0])*(max[1]-min[1])) * 0.5
}

func main() {
	p := [][]int{{4, 6}, {6, 5}, {3, 1}}
	fmt.Println(largestTriangleArea(p))
}
