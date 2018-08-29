package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	h := len(board)
	if h <= 0 {
		return false
	}
	// w := h
	conflictsX := make([][]byte, h)
	conflictsY := make([][]byte, h)
	for i := 0; i < h; i++ {
		for j := 0; j < h; j++ {
			if board[i][j] != '.' {
				continue
			}
			for m := i; m < h; m++ {
				if board[m][j] != '.' {
					conflictsX[m] = append(conflictsY[m], board[m][j])
				}
				for n := j; n < h; n++ {
					if board[m][n] != '.' {
						conflictsY[n] = append(conflictsY[n], board[m][n])
					}
				}
			}
		}
	}
	fmt.Println(conflictsX, conflictsY)
	return false
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
	isValidSudoku(sudoku)
}
