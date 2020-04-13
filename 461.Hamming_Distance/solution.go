package main

func hammingDistance(x int, y int) int {
	xor := x ^ y
	ret := 0
	for xor != 0 {
		if xor&0x01 == 1 {
			ret++
		}
		xor >>= 1
	}

	return ret
}
