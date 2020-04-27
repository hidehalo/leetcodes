package main

import (
	"fmt"
)

func numSpecialEquivGroups(A []string) int {
	counts := make([]int, 20)
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if len(A[i]) == len(A[j]) {
				for k := 0; k < len(A[i])-2; k++ {
					if A[i][k] == A[j][k+2] && A[i][k+2] == A[j][k] {
						counts[i]++
						counts[j]++
						fmt.Printf("%s special-equivalent %s\n", A[i], A[j])
						break
					}
				}
			}
		}
	}
	// sort.Ints(counts)
	fmt.Println(counts)

	return counts[19] + 1
}
