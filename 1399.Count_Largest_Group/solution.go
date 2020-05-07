package main

func bitCount(n int) int {
	ret := 0
	for n > 0 {
		if n&0x01 == 1 {
			ret++
		}
		n >>= 1
	}
	return ret
}

func countLargestGroup(n int) int {
	bitCountMap := make([]int, 14)
	for i := 1; i <= n; i++ {
		bitCountMap[i]++
	}
}
