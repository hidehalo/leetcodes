package main

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func sortedSquares(A []int) []int {
	counts := make(map[int]int)
	min := 10001
	max := -10001
	for i := 0; i < len(A); i++ {
		absVal := abs(A[i])
		if absVal > max {
			max = absVal
		}
		if absVal < min {
			min = absVal
		}
		counts[absVal]++
	}
	j := 0
	for i := min; i <= max; i++ {
		if v, ok := counts[i]; ok {
			for v > 0 {
				A[j] = i * i
				j++
				v--
			}
			continue
		}
	}

	return A
}
