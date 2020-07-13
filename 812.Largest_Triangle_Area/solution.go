package main

func largestTriangleArea(points [][]int) float64 {
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for i := 0; i < len(points); i++ {
		if i == 0 {
			minX, maxX, minY, maxY = points[0][0], points[0][0], points[0][1], points[0][1]
		} else {
			if points[i][0] < minX {
				minX = points[i][0]
			} else if points[i][0] > maxX {
				maxX = points[i][0]
			}
			if points[i][1] < minY {
				minY = points[i][1]
			} else if points[i][1] > maxY {
				maxY = points[i][1]
			}
		}
	}
}
