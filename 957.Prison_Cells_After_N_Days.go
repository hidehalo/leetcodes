package main

import (
	"fmt"
	"strconv"
)

// There are 8 prison cells in a row, and each cell is either occupied or vacant.

// Each day, whether the cell is occupied or vacant changes according to the following rules:

// If a cell has two adjacent neighbors that are both occupied or both vacant, then the cell becomes occupied.
// Otherwise, it becomes vacant.
// (Note that because the prison is a row, the first and the last cells in the row can't have two adjacent neighbors.)

// We describe the current state of the prison in the following way: cells[i] == 1 if the i-th cell is occupied, else cells[i] == 0.

// Given the initial state of the prison, return the state of the prison after N days (and N such changes described above.)

func bitmapToDecimal(bitmap []int) int {
	ret := 0
	size := len(bitmap)
	for i := 0; i < size; i++ {
		ret += (bitmap[i] << uint(size-i-1))
	}

	return ret
}

func decimalToBitmap(dec int) []int {
	ret := make([]int, 0, 1000)
	binStr := strconv.FormatInt(int64(dec), 2)
	for i := 0; i < len(binStr); i++ {
		ret = append(ret, int(binStr[i]-'0'))
	}

	return ret
}

func helper(cells []int, dp map[int]int) []int {
	if ret, ok := dp[bitmapToDecimal(cells)]; ok == true {
		return decimalToBitmap(ret)
	}
	key := bitmapToDecimal(cells)
	// FIXME: wrong algo
	for pivot := 1; pivot <= 6; pivot++ {
		if cells[pivot-1] == cells[pivot+1] {
			cells[pivot] = 1
		} else {
			cells[pivot] = 0
		}
		if pivot == 1 {
			cells[0], cells[7] = 0, 0
		}
	}
	dp[key] = bitmapToDecimal(cells)

	return cells
}

func prisonAfterNDays(cells []int, N int) []int {
	dp := make(map[int]int)
	for N > 0 {
		cells = helper(cells, dp)
		N--
	}

	return cells
}

func main() {
	cells := []int{1, 0, 0, 1, 0, 0, 1, 0}
	N := 1
	fmt.Println(prisonAfterNDays(cells, N))
	// excepted [0,1,1,1,1,0,0,0]
}
