package main

func numJewelsInStones(J string, S string) int {
	m := len(J)
	n := len(S)
	ret := 0
	types := make(map[byte]int)

	for i := 0; i < m; i++ {
		types[J[i]] = 1
	}

	for j := 0; j < n; j++ {
		if _, ok := types[S[j]]; ok {
			ret++
		}
	}

	return ret
}
