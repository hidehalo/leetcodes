package main

func getRow(rowIndex int) []int {
	prevCache := make([]int, rowIndex+1)
	ret := make([]int, rowIndex+1)
	for h := 0; h <= rowIndex; h++ {
		if h == 0 {
			ret[0] = 1
			prevCache[0] = 1
			continue
		}
		i, j := 0, h
		for i <= j {
			if i == 0 {
				ret[i], ret[j] = 1, 1
			} else {
				ret[i], ret[j] = prevCache[i]+prevCache[i-1], prevCache[i]+prevCache[i-1]
			}
			i++
			j--
		}
		copy(prevCache, ret)
	}

	return ret
}
