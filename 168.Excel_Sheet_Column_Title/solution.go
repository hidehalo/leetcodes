package main

var alphabets []byte

func convertToTitle(n int) string {
	ret := make([]byte, 0)
	if alphabets == nil {
		alphabets = make([]byte, 0)
		alphabets = append(alphabets, '@')
		for i := 0; i < 26; i++ {
			alphabets = append(alphabets, 'A'+byte(i))
		}
	}
	for n > 26 {
		key := n % 26
		if key > 0 {
			ret = append(ret, alphabets[key])
			n /= 26
		} else {
			n = n/26 - 1
			if n > 0 {
				ret = append(ret, 'Z')
			}
		}
	}
	ret = append(ret, alphabets[n])
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return string(ret)
}
