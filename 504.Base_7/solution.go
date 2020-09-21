package main

func convertToBase7(num int) string {
	var negative bool
	if num < 0 {
		num = -num
		negative = true
	}
	var ret []byte
	for {
		ret = append(ret, '0'+byte(num%7))
		num /= 7
		if num == 0 {
			break
		}
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	if negative {
		return "-" + string(ret)
	}
	return string(ret)
}
