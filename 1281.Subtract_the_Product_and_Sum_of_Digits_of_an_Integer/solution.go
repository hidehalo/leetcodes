package main

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for n >= 10 {
		tmp := n % 10
		sum += tmp
		product *= tmp
		n /= 10
	}
	sum += n
	product *= n

	return product - sum
}
