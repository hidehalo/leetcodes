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
	for idx = 0; idx < height; idx++ {
		_, ok = bsh(&matrix, idx, length, target)
		if ok {
			return true
		}
	}

	return false
}

func bsh(matrix *[][]int, cursor int, length int, target int) (int, bool) {
	if length == 1 {
		return 0, target == (*matrix)[cursor][0]
	}
	left := 0
	right := length - 1
	mid := (int)(right / 2)
	for left <= right {
		fmt.Println("h:", cursor, mid)
		if (*matrix)[cursor][mid] > target {
			right = mid - 1
			mid = (int)((right + left) / 2)
			continue
		} else if (*matrix)[cursor][mid] < target {
			left = mid + 1
			mid = (int)((right + left) / 2)
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
	right = height - 1
	mid = (int)(right / 2)
	for left <= right {
		fmt.Println("v:", mid, cursor)
		if (*matrix)[mid][cursor] > target {
			right = mid - 1
			mid = (int)((right + left) / 2)
			continue
		} else if (*matrix)[mid][cursor] < target {
			left = mid + 1
			mid = (int)((right + left) / 2)
			continue
		}

		return mid, true
	}

	return mid, false
}

func main() {
	nums := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	// nums := [][]int{
	// 	{1, 4},
	// 	{2, 5},
	// }
	// nums := [][]int{
	// 	{1, 1},
	// }
	// nums := [][]int{
	// 	{1, 2, 3, 4, 5},
	// 	{6, 7, 8, 9, 10},
	// 	{11, 12, 13, 14, 15},
	// 	{16, 17, 18, 19, 20},
	// 	{21, 22, 23, 24, 25},
	// }
	fmt.Println(searchMatrix(nums, 5))
}
