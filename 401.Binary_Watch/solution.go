package main

import "fmt"

func readBinaryWatch(num int) []string {
	if num >= 9 {
		return []string{}
	} else if num == 0 {
		return []string{"0:00"}
	}
	memoH := make([][]int, 4)
	memoM := make([][]int, 6)
	memoH[0] = []int{0}
	memoH[1] = []int{
		1, 2, 4, 8,
	}
	memoH[2] = []int{
		3, 5, 9, 6, 10,
	}
	memoH[3] = []int{
		7, 11,
	}
	memoM[0] = []int{0}
	memoM[1] = []int{
		1, 2, 4, 8, 16, 32,
	}
	memoM[2] = []int{
		3, 5, 9, 17, 33, 6, 10, 18, 34, 12, 20, 36, 24, 40, 48,
	}
	memoM[3] = []int{
		7, 11, 19, 35, 13, 21, 37, 25, 41, 49, 14, 22, 38, 26, 42, 50, 28, 44, 52, 56,
	}
	memoM[4] = []int{
		15, 23, 39, 27, 43, 51, 29, 45, 53, 57, 30, 46, 54, 58,
	}
	memoM[5] = []int{
		31, 45,
	}
	ret := make([]string, 0)
	for pivot := 0; pivot <= num; pivot++ {
		if pivot > 3 || num-pivot > 5 {
			continue
		}
		for _, h := range memoH[pivot] {
			for _, m := range memoM[num] {
				if m < 10 {
					ret = append(ret, fmt.Sprintf("%d:0%d", h, m))
				} else {
					ret = append(ret, fmt.Sprintf("%d:%d", h, m))
				}
			}
		}
	}
	return ret
}
