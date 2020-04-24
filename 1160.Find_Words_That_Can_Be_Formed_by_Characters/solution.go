package main

// slice is better than hash table
func countCharacters(words []string, chars string) int {
	buckets := make(map[byte]int)
	for i := 0; i < len(chars); i++ {
		buckets[chars[i]]++
	}
	ret := 0
	for i := 0; i < len(words); i++ {
		test := make(map[byte]int)
		for k, v := range buckets {
			test[k] = v
		}
		pass := 1
		for j := 0; j < len(words[i]); j++ {
			if test[words[i][j]] <= 0 {
				pass = 0
				break
			}
			test[words[i][j]]--
		}
		if pass == 1 {
			ret += len(words[i])
		}
	}

	return ret
}
