package main

func addDigits(num int) int {
	ret := 0
	for num > 0 {
		ret += num % 10
		num /= 10
	}
	if ret < 10 {
		return ret
	}
	return addDigits(ret)
}
