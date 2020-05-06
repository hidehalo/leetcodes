package main

func removeOuterParentheses(S string) string {
	ret := make([]byte, 0)
	l, r := 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			l++
			if l == 1 {
				continue
			}
		} else if S[i] == ')' {
			r++
		}
		if r > l { // illegal pattern
			return ""
		} else if l == r {
			l, r = 0, 0
			continue
		}
		ret = append(ret, S[i])
	}

	return string(ret)
}
