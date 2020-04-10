package main

func luckyNumbers (matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	min := make([]int,m)
	max := make([]int,n)
	for i:=0;i<m;i++ {
		min[i] = 100000
	}
	for j:=0;j<n;j++{
		max[j] = -1
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++{
			if matrix[i][j] < min[i] {
				min[i] = matrix[i][j]
			}
			if matrix[i][j] > max[j] {
				max[j] = matrix[i][j]
			}
		}
	}
	ret := make([]int,0)
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++{
			if min[i] == max[j] {
				ret = append(ret, min[i])
			}
		}
	}
	return ret
}
