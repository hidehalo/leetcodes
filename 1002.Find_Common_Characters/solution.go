package main

func commonChars(A []string) []string {
	buckets := make([]map[byte]int, len(A))
	ret := make([]string, 0)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if buckets[i] == nil {
				buckets[i] = make(map[byte]int)
			}
			buckets[i][A[i][j]]++
		}
	}

	uniq := make(map[byte]int)
	for key, val := range buckets[len(A)-1] {
		uniq[key] = val
	}
	for k := 0; k < len(buckets)-1; k++ {
		for key, val := range uniq {
			if buckets[k][key] < val {
				uniq[key] = buckets[k][key]
			}
		}
	}
	for k, v := range uniq {
		for v > 0 {
			ret = append(ret, string(k))
			v--
		}
	}

	return ret
}
