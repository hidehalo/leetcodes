package main

func reverseStr(s string, k int) string {
	offset := 0
	ret := []byte(s)
	for ; offset < len(ret); offset += 2 * k {
		needReverse := k
		if offset+k >= len(ret) {
			needReverse = len(ret) - offset
		}
		for i := 0; i < needReverse/2; i++ {
			ret[offset+i], ret[offset+needReverse-i-1] = ret[offset+needReverse-i-1], ret[offset+i]
		}
	}
	return string(ret)
}
