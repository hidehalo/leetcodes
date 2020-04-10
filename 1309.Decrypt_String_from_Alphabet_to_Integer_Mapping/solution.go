package main

import "fmt"

func freqAlphabets(s string) string {
	ret := make([]byte, 0)
	buf := [2]byte{}
	p := 0
	for i := 0; i < len(s); i++ {
		if p == 0 && s[i] > '2' {
			fmt.Printf(">> Decode %c to %c", s[i], 'a'+s[i]-'0')
			ret = append(ret, 'a'+s[i]-'0')
			continue
		}

		if s[i] == '#' {
			ret = append(ret, 'a'+10*(buf[0]-'0')+buf[1]-'0')
			p = 0
			continue
		}

		if p == 2 {
			for j := 0; j < p; j++ {
				ret = append(ret, 'a'+buf[j]-'0')
			}
			p = 0
		}
		buf[p] = s[i]
		p++
	}

	return string(ret)
}
