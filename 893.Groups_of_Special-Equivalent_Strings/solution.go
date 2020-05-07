package main

func numSpecialEquivGroups(A []string) int {
	hash := make(map[[52]int]struct{})
	for _, word := range A {
		var data [52]int
		for i, c := range word {
			// odd index alphabet number of S equal odd index alphabet number of T
			// S MUST special-equivalent T in odd index alphabets through some moves
			// the rule also works on even index alphabets
			if i%2 == 0 {
				data[c-'a']++
			} else {
				data[c-'a'+26]++
			}
		}
		hash[data] = struct{}{}
	}
	return len(hash)
}
