package main

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	bucket := make([][][]int, R+C-1)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			bucket[distance(r0, c0, i, j)] = append(bucket[distance(r0, c0, i, j)], []int{i, j})
		}
	}
	ret := make([][]int, 0)
	for i := 0; i < R+C-1; i++ {
		if len(bucket[i]) > 0 {
			ret = append(ret, bucket[i]...)
		}
	}
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(r0, c0, ri, ci int) int {
	return abs(r0-ri) + abs(c0-ci)
}
