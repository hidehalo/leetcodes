package main

func findWords(words []string) []string {
	lineMap := make(map[byte]int)
	for _, b := range []byte{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'} {
		lineMap[b] = 1
		lineMap[b-32] = 1
	}
	for _, b := range []byte{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'} {
		lineMap[b] = 2
		lineMap[b-32] = 2
	}
	for _, b := range []byte{'z', 'x', 'c', 'v', 'b', 'n', 'm'} {
		lineMap[b] = 3
		lineMap[b-32] = 3
	}
	ret := make([]string, 0)
	for i := 0; i < len(words); i++ {
		tested := true
		for j := 0; j < len(words[i])-1; j++ {
			if lineMap[words[i][j]] != lineMap[words[i][j+1]] {
				tested = false
				break
			}
		}
		if tested == true {
			ret = append(ret, words[i])
		}
	}
	return ret
}
