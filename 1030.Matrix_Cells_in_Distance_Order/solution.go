package main

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	ret := make([][]int, 0)
	for d := 0; d < R+C-1; d++ {
		if d == 0 {
			ret = append(ret, []int{r0, c0})
			continue
		}
		for k := 0; k <= d; k++ {
			if r0+k < R && c0+d-k >= 0 && c0+d-k < C {
				ret = append(ret, []int{r0 + k, c0 + d - k})
			}
		}
	}

	return ret
}
