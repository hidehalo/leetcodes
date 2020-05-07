package main

func islandPerimeter(grid [][]int) int {
	ret := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				if !inMap(len(grid), len(grid[i]), i+1, j) || grid[i+1][j] == 0 {
					ret++
				}
				if !inMap(len(grid), len(grid[i]), i-1, j) || grid[i-1][j] == 0 {
					ret++
				}
				if !inMap(len(grid), len(grid[i]), i, j+1) || grid[i][j+1] == 0 {
					ret++
				}
				if !inMap(len(grid), len(grid[i]), i, j-1) || grid[i][j-1] == 0 {
					ret++
				}
			}
		}
	}

	return ret
}

func inMap(M, N, i, j int) bool {
	if i >= 0 && i < M && j >= 0 && j < N {
		return true
	}

	return false
}
