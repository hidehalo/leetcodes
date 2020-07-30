package main

func hammingWeight(num uint32) int {
	ret := 0
	for num > 0 {
		if num&0x01 == 0x01 {
			ret++
		}
		num >>= 1
	}
	return ret
}
