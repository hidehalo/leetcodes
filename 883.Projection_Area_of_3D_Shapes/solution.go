package main

func projectionArea(grid [][]int) int {
	colMax := make([]int, len(grid[0]))
	rowMax := make([]int, len(grid))
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > rowMax[i] {
				rowMax[i] = grid[i][j]
			}
			if grid[i][j] > colMax[j] {
				colMax[j] = grid[i][j]
			}
			if grid[i][j] > 0 {
				ret++
			}
		}
	}
	for _, v := range colMax {
		ret += v
	}
	for _, v := range rowMax {
		ret += v
	}

	return ret
}
