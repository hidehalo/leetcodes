package main

func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}
	occur := make(map[int]int)
	for _, v := range counts {
		occur[v]++
		if occur[v] > 1 {
			return false
		}
	}

	return true
}
