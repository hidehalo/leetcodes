package main

func diStringMatch(S string) []int {
	ret := make([]int, 0)
	min := 0
	max := len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == 'I' {
			ret = append(ret, min)
			min++
		} else if S[i] == 'D' {
			ret = append(ret, max)
			max--
		}
	}
	for min <= max {
		ret = append(ret, min)
		min++
	}
	return ret
}
