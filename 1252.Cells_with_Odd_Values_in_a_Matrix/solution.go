package main

func oddCells(n int, m int, indices [][]int) int {
	rowCount := make([]int, n)
	colCount := make([]int, m)
	for _, index := range indices {
		rowCount[index[0]]++
		colCount[index[1]]++
	}
	oddRow, evenRow := 0, 0
	oddCol, evenCol := 0, 0
	for _, c := range rowCount {
		if c%2 != 0 {
			oddRow++
		} else {
			evenRow++
		}
	}
	for _, c := range colCount {
		if c%2 != 0 {
			oddCol++
		} else {
			evenCol++
		}
	}

	return oddRow*(m-oddCol) + oddCol*(n-oddRow)
}
