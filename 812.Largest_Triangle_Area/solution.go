package main

import "math"

// func largestTriangleArea(points [][]int) float64 {
// 	minX, maxX, minY, maxY := 0, 0, 0, 0
// 	for i := 0; i < len(points); i++ {
// 		if i == 0 {
// 			minX, maxX, minY, maxY = points[0][0], points[0][0], points[0][1], points[0][1]
// 		} else {
// 			if points[i][0] < minX {
// 				minX = points[i][0]
// 			} else if points[i][0] > maxX {
// 				maxX = points[i][0]
// 			}
// 			if points[i][1] < minY {
// 				minY = points[i][1]
// 			} else if points[i][1] > maxY {
// 				maxY = points[i][1]
// 			}
// 		}
// 	}
// 	can := make([]int, 0)
// 	for i := 0; i < len(points); i++ {
// 		if points[i][0] == minX || points[i][0] == maxX || points[i][1] == minY || points[i][1] == maxY {
// 			can = append(can, i)
// 		}
// 	}
// 	maxS := float64(0)
// 	for i := 0; i < len(can); i++ {
// 		for j := i + 1; j < len(can); j++ {
// 			for k := j + 1; k < len(can); k++ {
// 				s := float64((points[i][0] - points[j][0]) * (points[j][1] - points[k][1]) / 2)
// 				if s < 0 {
// 					s = -s
// 				}
// 				if s > maxS {
// 					maxS = s
// 				}
// 			}
// 		}
// 	}

// 	return maxS
// }

func largestTriangleArea(points [][]int) float64 {
	size := len(points)
	var ret float64
	for i := 0; i < size-2; i++ {
		for j := i + 1; j < size-1; j++ {
			for k := 0; k < size; k++ {
				if v := cal(points[i][0], points[i][1], points[j][0], points[j][1], points[k][0], points[k][1]); v > ret {
					ret = v
				}
			}
		}
	}
	return ret
}

func cal(x1, y1, x2, y2, x3, y3 int) float64 {
	return math.Abs(float64(x1*y2+x2*y3+x3*y1-y1*x2-y2*x3-y3*x1)) / 2.0
}
