package main

func shortestToChar(S string, C byte) []int {
	ranges := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			ranges = append(ranges, i)
		}
	}
	ret := make([]int, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == C {
			ret = append(ret, 0)
		} else {
			min := len(S) + 1
			for j := 0; j < len(ranges); j++ {
				if ranges[j] > i {
					if ranges[j]-i < min {
						min = ranges[j] - i
					}
					break
				}
				min = i - ranges[j]
			}
			ret = append(ret, min)
		}
	}

	return ret
}
