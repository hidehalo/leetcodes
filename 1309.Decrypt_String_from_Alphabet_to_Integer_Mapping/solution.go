package main

func freqAlphabets(s string) string {
	ret := make([]byte, 0)
	i, j, k := 0, 1, 2

	for i < len(s)-2 {
		if s[k] == '#' {
			ret = append(ret, 'a'+10*(s[i]-'0')+s[j]-'1')
			i = k + 1
			j = i + 1
			k = j + 1
		} else {
			ret = append(ret, 'a'+s[i]-'1')
			i++
			j++
			k++
		}
	}

	for ; i < len(s); i++ {
		ret = append(ret, 'a'+s[i]-'1')
	}

	return string(ret)
}
