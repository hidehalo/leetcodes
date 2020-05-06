package main

func destCity(paths [][]string) string {
	cityCount := make(map[string]int)
	tails := make([]string, 0)
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			if j == len(paths[i])-1 {
				tails = append(tails, paths[i][j])
			} else {
				cityCount[paths[i][j]]++
			}
		}
	}
	for _, tail := range tails {
		if cityCount[tail] == 0 {
			return tail
		}
	}
	return ""
}
