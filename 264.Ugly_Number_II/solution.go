package main

func nthUglyNumber(n int) int {
	hit := 0
	number := 1
	for hit < n {
		if isUgly(number) {
			hit++
		}
		number++
	}
	return number - 1
}

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num%2 == 0 {
		num /= 2
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num == 1
}
