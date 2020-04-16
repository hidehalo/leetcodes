package main

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ret := 0
	for i := 0; i < len(arr1); i++ {
		tested := false
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				tested = false
				break
			}
			tested = true
		}
		if tested == true {
			ret++
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
