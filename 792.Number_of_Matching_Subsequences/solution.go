package main

import "unsafe"

// Given a string s and an array of strings words, return the number of words[i] that is a subsequence of s.
// A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
// For example, "ace" is a subsequence of "abcde".

func numMatchingSubseq(s string, words []string) int {
	count := 0
	bs := []byte(s)
	sMeta := generateMeta(bs)
	wMeta := make([]int, 26)
	dp := make(map[string]int8)
	for i := 0; i < len(words); i++ {
		if _, has := dp[words[i]]; has {
			count++
			continue
		}
		bwp := (*[]byte)(unsafe.Pointer(&words[i]))
		calcMeta(*bwp, wMeta)
		if !checkMeta(sMeta, wMeta) {
			continue
		}
		if !isSubsequence(bs, *bwp) {
			continue
		}
		dp[words[i]] = int8(1)
		count++
	}
	return count
}

// T(n)
// S(26)
func generateMeta(s []byte) []int {
	meta := make([]int, 26)
	calcMeta(s, meta)
	return meta
}

// T(n)
func calcMeta(s []byte, meta []int) {
	for i := 0; i < len(meta); i++ {
		meta[i] = 0
	}
	for i := 0; i < len(s); i++ {
		meta[int(s[i]-'a')]++
	}
}

// T(26)
func checkMeta(meta1, meta2 []int) bool {
	for i := 0; i < 26; i++ {
		if meta2[i] > meta1[i] {
			return false
		}
	}
	return true
}

// T(m*n)
func isSubsequence(t, s []byte) bool {
	if len(s) > len(t) {
		return false
	} else if len(s) == 0 {
		return true
	}
	var p, q int
	for p < len(s) {
		for q < len(t) && t[q] != s[p] {
			q++
		}
		if q == len(t) {
			return false
		} else if t[q] == s[p] {
			q++
		}
		p++
	}

	return true
}
