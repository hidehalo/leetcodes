package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	var idx int
	var ok bool
	height := len(matrix)
	if height <= 0 {
		return false
	}
	length := len(matrix[0])
	if length <= 0 {
		return false
	}
	idx, ok = bsh(&matrix, 0, length, target)
	if ok != true {
		idx, ok = bsv(&matrix, idx, height, target)
	}
	if ok != true {
		_, ok = bsh(&matrix, idx, length, target)

		return ok
	}

	return true
}

func bsh(matrix *[][]int, cursor int, length int, target int) (int, bool) {
	if length == 1 {
		return 0, target == (*matrix)[0][0]
	}
	left := 0
	right := length
	mid := (int)(length / 2)
	for mid >= left && mid < right {
		// fmt.Println(cursor, mid)
		if (*matrix)[cursor][mid] > target {
			if left+1 == right {
				break
			}
			right = mid
			mid = (int)((left + right) / 2)
			continue
		} else if (*matrix)[cursor][mid] < target {
			if left+1 == right {
				break
			}
			left = mid
			mid = (int)((left + right) / 2)
			continue
		}

		return mid, true
	}

	return mid, false
}

func bsv(matrix *[][]int, cursor int, height int, target int) (int, bool) {
	if height == 1 {
		return 0, target == (*matrix)[0][cursor]
	}
	var mid, left, right int
	left = 0
	right = height
	mid = (int)(height / 2)
	for mid >= left && mid < right {
		// fmt.Println(mid, cursor)
		if (*matrix)[mid][cursor] > target {
			if left+1 == right {
				break
			}
			right = mid
			mid = (int)((left + right) / 2)
			continue
		} else if (*matrix)[mid][cursor] < target {
			if left+1 == right {
				break
			}
			left = mid
			mid = (int)((left + right) / 2)
			continue
		}

		return mid, true
	}

	return mid, false
}

func main() {
	nums := [][]int{
		{1, 4},
		{2, 3},
	}
	fmt.Println(searchMatrix(nums, 5))
}
