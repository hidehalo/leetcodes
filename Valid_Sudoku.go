package main

import (
	"fmt"
	"math"
)

func isValidSudoku(board [][]byte) bool {
	h := len(board)
	if h <= 0 {
		return false
	}
	// w := h
	conflictsX := make([]map[byte]bool, h)
	conflictsY := make([]map[byte]bool, h)
	conflictsArea := make([]map[byte]bool, h)
	for k := 0; k < h; k++ {
		conflictsX[k] = make(map[byte]bool)
		conflictsY[k] = make(map[byte]bool)
		conflictsArea[k] = make(map[byte]bool)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < h; j++ {
			if board[i][j] != '.' {
				if conflictsY[i][board[i][j]] == true || conflictsX[j][board[i][j]] == true || conflictsArea[area(i, j, h)][board[i][j]] == true {
					return false
				}
				conflictsY[i][board[i][j]] = true
				conflictsX[j][board[i][j]] = true
				conflictsArea[area(i, j, h)][board[i][j]] = true
			}
		}
	}

	return true
}

func area(x int, y int, h int) int {
	sqrth := int(math.Sqrt(float64(h)))

	return y/sqrth + sqrth*(x/sqrth)
}

func main() {
	sudoku := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(sudoku))
}
