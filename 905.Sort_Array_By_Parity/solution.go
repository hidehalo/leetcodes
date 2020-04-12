package main

func sortArrayByParity(A []int) []int {
	j := len(A) - 1
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			continue
		}
		for i < j {
			if A[j]%2 == 0 {
				A[i] ^= A[j]
				A[j] ^= A[i]
				A[i] ^= A[j]
				j--
				break
			}
			j--
		}
	}

	return A
}
