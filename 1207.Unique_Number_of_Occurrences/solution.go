package main

func uniqueOccurrences(arr []int) bool {
	counts := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if counts[arr[i]] > 0 {
			return false
		}
		counts[arr[i]]++
	}

	return true
}
