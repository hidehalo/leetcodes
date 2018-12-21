package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	buffer := make([]byte, 0, len(num1)+len(num2))
	carray := 0
	char := byte('0')
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			char, carray = helper(num1[i], num2[j], carray)
			buffer = append(buffer, char)
		}
	}

	return string(buffer)
}

func helper(a, b byte, carry int) (byte, int) {
	ret := int(a-'0')*int(b-'0') + carry

	return byte('0' + ret%10), int(ret / 10)
}

func main() {
	num1 := "120"
	num2 := "11"
	fmt.Println(multiply(num1, num2))
}
