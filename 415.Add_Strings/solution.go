package main

func addStrings(num1 string, num2 string) string {
	var carry, val int
	ret := make([]byte, 0)
	for i := 1; ; {
		if len(num1)-i < 0 && len(num2)-i < 0 {
			if carry > 0 {
				ret = append(ret, '1')
			}
			break
		}
		if len(num1)-i >= 0 && len(num2)-i >= 0 {
			val = int(num1[len(num1)-i]+num2[len(num2)-i]-'0'-'0') + carry
		} else if len(num1)-i >= 0 {
			val = int(num1[len(num1)-i]-'0') + carry
		} else if len(num2)-i >= 0 {
			val = int(num2[len(num2)-i]-'0') + carry
		}
		carry = 0
		if val > 9 {
			val -= 10
			carry++
		}
		ret = append(ret, byte(val)+'0')
		i++
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return string(ret)
}
