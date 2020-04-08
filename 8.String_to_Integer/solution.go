package main

import (
	"../ds"
)

func check(b byte) int {
	if b >= '0' && b <= '9' {
		return 1
	} else if b == '-' {
		return -1
	} else if b == '+' {
		return -2
	}

	return 0
}

func myAtoi(str string) int {
	size := len(str)
	stack := ds.NewStack()
	flag := 1
	offset := 0
	for ; offset < size; offset++ {
		if str[offset] == ' ' {
			continue
		} else {
			break
		}
	}
	sign := 0
	for ; offset < size; offset++ {
		if str[offset] == '+' || str[offset] == '-' {
			sign++
			if sign > 1 {
				return 0
			}
			if str[offset] == '+' {
				flag = 1
			} else {
				flag = -1
			}
			continue
		} else {
			break
		}
	}
	for ; offset < size; offset++ {
		if str[offset] == '0' {
			continue
		} else {
			break
		}
	}
	for i := offset; i < size; i++ {
		check := check(str[i])
		if i == offset && check == 0 {
			return 0
		} else if check != 1 && !stack.Empty() {
			break
		} else if stack.Len() >= 10 {
			if flag == 1 {
				return 2147483647
			}
			return -2147483648
		}
		switch check {
		case 0:
			break
		case -1:
			return 0
		case -2:
			return 0
		case 1:
			stack.Push(str[i])
		}
	}
	ret := 0
	step := 1
	for !stack.Empty() {
		n := int(stack.Pop()) - 48
		ret += n * step
		step *= 10
	}

	if ret*flag > 2147483647 {
		return 2147483647
	} else if ret*flag < -2147483648 {
		return -2147483648
	}

	return ret * flag
}
