package main

func toHex(num int) string {
	if num < 0 {
		num = 0xffffffff + 1 + num
	} else if num == 0 {
		return "0"
	}
	var remainder int
	ret := make([]byte, 0)

	for num > 0 {
		remainder = num % 16
		num /= 16
		if remainder < 10 {
			ret = append(ret, '0'+byte(remainder))
		} else {
			ret = append(ret, 'a'+byte(remainder-10))
		}
	}

	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}

	return string(ret)
}
