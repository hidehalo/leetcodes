package main

func canConstruct(ransomNote string, magazine string) bool {
	repo := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		repo[magazine[i]]++
	}
	for j := 0; j < len(ransomNote); j++ {
		if repo[ransomNote[j]] <= 0 {
			return false
		}
		repo[ransomNote[j]]--
	}
	return true
}
