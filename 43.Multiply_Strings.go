package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	buffer := make([]byte, len(num1)+len(num2))
	for k := 0; k < len(buffer); k++ {
		buffer[k] = '0'
	}
	round := 0
	for i := len(num1) - 1; i >= 0; i-- {
		offset := len(buffer) - 1 - round
		carray := 0
		addCarray := 0
		char := byte('0')
		for j := len(num2) - 1; j >= 0; j-- {
			fmt.Printf("执行%c * %c + %d ", num1[i], num2[j], carray)
			char, carray = product(num1[i], num2[j], carray)
			fmt.Printf("= %c 进 %d\n", char, carray)
			fmt.Printf("执行%c + %c + %d ", buffer[offset], char, addCarray)
			buffer[offset], addCarray = add(buffer[offset], char, addCarray)
			fmt.Printf("= %c 进 %d\n", buffer[offset], addCarray)
			offset--
		}
		round++
		buffer[offset] = byte(carray + addCarray + '0')
		fmt.Println(string(buffer))
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

func main() {
	num1 := "23312"
	num2 := "7418"
	fmt.Println(multiply(num1, num2))
}
