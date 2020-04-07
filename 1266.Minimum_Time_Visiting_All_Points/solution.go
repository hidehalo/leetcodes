package main

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func minTimeToVisitAllPoints(points [][]int) int {
	ret := 0
	for i := 0; i < len(points)-1; i++ {
		s, f := points[i], points[i+1]
		dx := abs(f[0] - s[0])
		dy := abs(f[1] - s[1])
		if dx > dy {
			ret += dx
		} else {
			ret += dy
		}
	}

	return ret
}
