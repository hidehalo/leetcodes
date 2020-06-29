package main

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	i := 0
	ret := make([]byte, 0)
	needCarry := false
	ai, bi := byte('0'), byte('0')
	for i < m || i < n {
		if i < m {
			ai = a[m-1-i]
		} else {
			ai = '0'
		}
		if i < n {
			bi = b[n-1-i]
		} else {
			bi = '0'
		}
		sum := byte('0')
		if ai == '1' && bi == '1' {
			if needCarry == true {
				sum = '1'
			}
			needCarry = true
		} else if ai == '0' && bi == '0' {
			if needCarry == true {
				sum = '1'
			}
			needCarry = false
		} else {
			if needCarry == true {
				sum = '0'
			} else {
				needCarry = false
				sum = '1'
			}
		}
		ret = append(ret, sum)
		i++
	}
	if needCarry == true {
		ret = append(ret, byte('1'))
	}
	i, j := 0, len(ret)-1
	for i < j {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}

	return string(ret)
}
