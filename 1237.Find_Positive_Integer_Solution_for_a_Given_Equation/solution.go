package main

func findSolution(customFunction func(int, int) int, z int) [][]int {
	ret := make([][]int, 0)
	var i, j int
	for i = 1; ; i++ {
		for j = 1; ; j++ {
			if customFunction(i, j) > z {
				break
			} else if customFunction(i, j) == z {
				ret = append(ret, []int{i, j})
			}
		}
		if j == 1 {
			break
		}
	}

	return ret
}
