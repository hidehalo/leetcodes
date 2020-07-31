package main

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}
	c := 1
	p := 1
	for c < n {
		c <<= p
	}

	return c == n
}
