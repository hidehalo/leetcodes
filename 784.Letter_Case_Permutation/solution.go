package main

func letterCasePermutation(S string) []string {
	ret := make([]string, 0)
	for i := 0; i < len(S); i++ {
		if 'a' <= S[i] && S[i] <= 'z' {
			upper := S[i] - 32
			newStr := []byte(S)
			for _, rs := range ret {
				renewStr := []byte(rs)
				renewStr[i] = upper
				ret = append(ret, string(renewStr))
			}
			newStr[i] = upper
			ret = append(ret, string(newStr))
		} else if 'A' <= S[i] && S[i] <= 'Z' {
			lower := 32 + S[i]
			newStr := []byte(S)
			for _, rs := range ret {
				renewStr := []byte(rs)
				renewStr[i] = lower
				ret = append(ret, string(renewStr))
			}
			newStr[i] = lower
			ret = append(ret, string(newStr))
		}
	}
	ret = append(ret, S)

	return ret
}
