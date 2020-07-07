package main

func plusOne(digits []int) []int {
	needCarry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if needCarry == true {
			digits[i]++
		}
		if digits[i] == 10 {
			needCarry = true
			digits[i] = 0
		} else {
			needCarry = false
		}
	}
	if needCarry {
		return append([]int{1}, digits...)
	}
	return digits
}
