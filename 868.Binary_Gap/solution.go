package main

func binaryGap(N int) int {
	if N <= 2 {
		return 0
	}
	positaveBits := make([]int, 0)
	bitNo := 1
	for N > 0 {
		if N&0x01 == 1 {
			positaveBits = append(positaveBits, bitNo)
		}
		N >>= 1
		bitNo++
	}
	ret := 0
	for i := len(positaveBits) - 1; i > 0; i-- {
		j := i - 1
		if positaveBits[i]-positaveBits[j] > ret {
			ret = positaveBits[i] - positaveBits[j]
		}
	}

	return ret
}
