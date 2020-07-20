package main

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for h := 0; h < numRows; h++ {
		ret[h] = make([]int, h+1)
		if h == 0 {
			ret[h][0] = 1
			continue
		}
		i, j := 0, h
		for i <= j {
			if i == 0 {
				ret[h][i], ret[h][j] = 1, 1
			} else {
				ret[h][i], ret[h][j] = ret[h-1][i]+ret[h-1][i-1], ret[h-1][i]+ret[h-1][i-1]
			}
			i++
			j--
		}
	}

	return ret
}
