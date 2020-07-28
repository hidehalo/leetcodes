package main

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	weight := uint32(1 << 31)
	for num > 0 {
		if num&0x01 == 1 {
			ret += weight
		}
		num >>= 1
		weight >>= 1
	}
	return ret
}
