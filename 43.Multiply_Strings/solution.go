package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	maxSize := len(num1) + len(num2)
	buffer := make([]byte, maxSize)
	for k := 0; k < len(buffer); k++ {
		buffer[k] = '0'
	}
	round := 0
	for i := len(num1) - 1; i >= 0; i-- {
		offset := maxSize - 1 - round
		carray := 0
		addCarray := 0
		char := byte('0')
		for j := len(num2) - 1; j >= 0; j-- {
			char, carray = product(num1[i], num2[j], carray)
			buffer[offset], addCarray = add(buffer[offset], char, addCarray)
			offset--
		}
		round++
		buffer[offset] = byte(carray + addCarray + '0')
	}
	if buffer[0] == '0' {
		buffer = buffer[1:]
	}

	return string(buffer)
}

func product(a, b byte, carry int) (byte, int) {
	ret := int(a-'0')*int(b-'0') + carry

	return byte('0' + ret%10), int(ret / 10)
}

func add(a, b byte, carry int) (byte, int) {
	ret := int(a-'0') + int(b-'0') + carry

	return byte('0' + ret%10), int(ret / 10)
}
