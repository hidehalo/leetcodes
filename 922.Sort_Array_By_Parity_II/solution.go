package main

func sortArrayByParityII(A []int) []int {
	p, q := 0, 1
	for q < len(A) {
		if A[p]%2 == 0 {
			p += 2
		}
		if A[q]%2 != 0 {
			q += 2
		}
		if A[q]%2 == 0 && A[p]%2 != 0 {
			A[q], A[p] = A[p], A[q]
		}
	}

	return A
}
