package main

func findNumbers(nums []int) int {
	ret := 0
	for _, v := range nums {
		if isEvenNumberOfDigits(v) {
			ret++
		}
	}

	return ret
}

func isEvenNumberOfDigits(n int) bool {
	if n >= 1e4 {
		return false
	} else if n >= 1e3 {
		return true
	} else if n >= 1e2 {
		return false
	} else if n >= 1e1 {
		return true
	}

	return false
}
