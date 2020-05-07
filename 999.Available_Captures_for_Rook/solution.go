package main

func numRookCaptures(board [][]byte) int {
	var rx, ry int
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				rx, ry = i, j
				break
			}
		}
	}
	ret := 0
	for i := rx - 1; i >= 0; i-- {
		if board[i][ry] == 'p' {
			ret++
			break
		} else if board[i][ry] == 'B' {
			break
		}
	}
	for i := rx + 1; i < len(board); i++ {
		if board[i][ry] == 'p' {
			ret++
			break
		} else if board[i][ry] == 'B' {
			break
		}
	}
	for i := ry - 1; i >= 0; i-- {
		if board[rx][i] == 'p' {
			ret++
			break
		} else if board[rx][i] == 'B' {
			break
		}
	}
	for i := ry + 1; i < len(board[rx]); i++ {
		if board[rx][i] == 'p' {
			ret++
			break
		} else if board[rx][i] == 'B' {
			break
		}
	}

	return ret
}
