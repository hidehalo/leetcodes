package main

func numberOfBoomerangs(points [][]int) int {
	ret := 0
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(points); k++ {
				if k == j || k == i {
					continue
				}
				if pow2(points[i][0]-points[j][0])+pow2(points[i][1]-points[j][1]) == pow2(points[i][0]-points[k][0])+pow2(points[i][1]-points[k][1]) {
					ret++
				}
			}
		}
	}
	return ret
}

func pow2(a int) int {
	return a * a
}
