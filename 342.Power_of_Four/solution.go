package main

func isPowerOfFour(num int) bool {
	for num >= 4 {
		if num%4 != 0 {
			return false
		}
		num >>= 2
	}

	return num == 1
}
