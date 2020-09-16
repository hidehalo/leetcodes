package main

import "math"

func constructRectangle(area int) []int {
	sqrt2 := math.Sqrt(float64(area))
	var W int
	for W = int(sqrt2); W > 0; {
		if area%W == 0 {
			break
		}
		W--
	}
	return []int{area / W, W}
}
