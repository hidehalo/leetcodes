package main

import (
	"fmt"
	"math"
)

func solveSudoku(board [][]byte) bool {
	h := len(board)
	if h <= 0 {
		return
	}
	filled := 0
	for i := 0; i < h; i++ {
		for j := 0; j < h; j++ {
			if board[i][j] == '.' {
				for k := 1; k < h; k++ {
					board[i][j] = byte(k + '0')
					if valid(board, i, j) {
						break
					} else {
						board[i][j] = '.'
					}
				}
			} else {
				filled++
			}
		}
	}
}

func valid(board [][]byte, x int, y int) bool {
	h := len(board)
	if h <= 0 {
		return false
	}
	checkX := make([]int, h)
	checkY := make([]int, h)
	checkA := make([]int, h)
	sqrth := int(math.Sqrt(float64(h)))
	for i := 0; i < h; i++ {
		if board[x][i] != '.' {
			numX := int(board[x][i] - '1')
			checkX[numX]++
			if checkX[numX] > 1 {
				return false
			}
		}
		if board[i][y] != '.' {
			numY := int(board[i][y] - '1')
			checkY[numY]++
			if checkY[numY] > 1 {
				return false
			}
		}
		ax := area(x, y, h)%sqrth*sqrth + i%sqrth
		ay := area(x, y, h)/sqrth*sqrth + i/sqrth
		if board[ax][ay] != '.' {
			numA := int(board[ax][ay] - '1')
			checkA[numA]++
			if checkA[numA] > 1 {
				return false
			}
		}
	}

	return true
}

func area(x int, y int, h int) int {
	sqrth := int(math.Sqrt(float64(h)))

	return y/sqrth + sqrth*(x/sqrth)
}

func printh(board [][]byte) {
	h := len(board)
	if h <= 0 {
		return
	}
	for i := 0; i < h; i++ {
		for j := 0; j < h; j++ {
			fmt.Printf(" %s ", string(board[i][j]))
		}
		fmt.Println("")
	}
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
	printh(sudoku)
	fmt.Println()
	solveSudoku(sudoku)
	printh(sudoku)
}
