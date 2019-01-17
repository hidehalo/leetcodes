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
	ret := make([]int, 8)
	binStr := strconv.FormatInt(int64(dec), 2)
	for i := 8 - len(binStr); i < 8; i++ {
		ret[i] = int(binStr[i-8+len(binStr)] - '0')
	}

	return ret
}

func getCircleRound(dp map[int]int, start int) int {
	ret := 0
	next := start
	for {
		next = dp[next]
		ret++
		if next == start {
			break
		}
	}

	return ret
}

func procedure(cells []int, buffer []int, dp map[int]int) {
	if ret, ok := dp[bitmapToDecimal(cells)]; ok == true {
		buffer = decimalToBitmap(ret)
	} else {
		for pivot := 1; pivot <= 6; pivot++ {
			if cells[pivot-1] == cells[pivot+1] {
				buffer[pivot] = 1
			} else {
				buffer[pivot] = 0
			}
		}
	}
	key := bitmapToDecimal(cells)
	dp[key] = bitmapToDecimal(buffer)
	for i, v := range buffer {
		cells[i] = v
	}
}

func prisonAfterNDays(cells []int, N int) []int {
	dp := make(map[int]int)
	buffer := make([]int, 8)
	var start, current int
	tmp := N
	circleRound := 0
	for tmp > 0 {
		current = bitmapToDecimal(cells)
		if _, exsit := dp[current]; exsit == true {
			start = current
			circleRound = getCircleRound(dp, start)
			break
		}
		procedure(cells, buffer, dp)
		tmp--
	}
	if circleRound > 0 {
		offset := len(dp) - circleRound
		remain := (N - offset) % circleRound
		key := start
		for i := 0; i < remain; i++ {
			key = dp[key]
		}

		return decimalToBitmap(key)
	}

	return cells
}

func main() {
	cells := []int{1, 0, 0, 1, 0, 0, 1, 0}
	N := 1000000000
	fmt.Println(prisonAfterNDays(cells, N))
	//e : 0,0,1,1,1,1,1,0
	cells = []int{0, 0, 1, 1, 1, 1, 0, 0}
	N = 8
	fmt.Println(prisonAfterNDays(cells, N))
	//e : [0,0,0,1,1,0,0,0]
	cells = []int{0, 1, 0, 1, 1, 0, 0, 1}
	N = 7
	fmt.Println(prisonAfterNDays(cells, N))
	//e : [0,0,1,1,0,0,0,0]
}
