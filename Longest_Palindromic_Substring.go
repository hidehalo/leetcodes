package main

import (
	"fmt"
)

/**
 * Brute Force
 * O(n) = n^3
 */
func longestPalindrome(s string) string {
	size := len(s)
	if size == 1 {
		return s
	}
	ret := ""
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			ti := i
			tj := j
			for ti < tj {
				if s[ti] != s[tj] {
					break
				}
				ti++
				tj--
			}
			if ti >= tj && len(ret) < j-i+1 {
				ret = s[i : j+1]
			}
		}
	}

	return ret
}

/**
 * DP Pseudocode
 * O(n^2)
 */
//  function LCSubstr(S[1..r], T[1..n])
//  L := array(1..r, 1..n)
//  z := 0
//  ret := {}
//  for i := 1..r
// 	 for j := 1..n
// 		 if S[i] == T[j]
// 			 if i == 1 or j == 1
// 				 L[i,j] := 1
// 			 else
// 				 L[i,j] := L[i-1,j-1] + 1
// 			 if L[i,j] > z
// 				 z := L[i,j]
// 				 ret := {S[i-z+1..i]}
// 			 else
// 			 if L[i,j] == z
// 				 ret := ret ∪ {S[i-z+1..i]}
// 		 else
// 			 L[i,j] := 0
//  return ret

func main() {
	fmt.Println(longestPalindrome("jhgtrclvzumufurdemsogfkpzcwgyepdwucnxrsubrxadnenhvjyglxnhowncsubvdtftccomjufwhjupcuuvelblcdnuchuppqpcujernplvmombpdttfjowcujvxknzbwmdedjydxvwykbbamfnsyzcozlixdgoliddoejurusnrcdbqkfdxsoxxzlhgyiprujvvwgqlzredkwahexewlnvqcwfyahjpeiucnhsdhnxtgizgpqphunlgikogmsffexaeftzhblpdxrxgsmeascmqngmwbotycbjmwrngemxpfakrwcdndanouyhnnrygvntrhcuxgvpgjafijlrewfhqrguwhdepwlxvrakyqgstoyruyzohlvvdhvqmzdsnbtlwctetwyrhhktkhhobsojiyuydknvtxmjewvssegrtmshxuvzcbrabntjqulxkjazrsgbpqnrsxqflvbvzywzetrmoydodrrhnhdzlajzvnkrcylkfmsdode"))
}
