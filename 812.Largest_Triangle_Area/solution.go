package main

import "math"

func areaOfTriangle(p1, p2, p3 []int) float64 {
	area := p1[0]*p2[1] +
		p2[0]*p3[1] +
		p3[0]*p1[1] -
		p1[1]*p2[0] -
		p2[1]*p3[0] -
		p3[1]*p1[0]
	return math.Abs(float64(area) / 2.0)
}

func largestTriangleArea(points [][]int) float64 {
	maxS := float64(0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				s := areaOfTriangle(points[i], points[j], points[k])
				if s > maxS {
					maxS = s
				}
			}
		}
	}

	return maxS
}
