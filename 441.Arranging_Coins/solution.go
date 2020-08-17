package main

func arrangeCoins(n int) int {
	k := 0
	kn := 0
	for {
		k++
		if kn+k > n {
			return k - 1
		} else if kn+k == n {
			return k
		}
		kn += k
	}
}
