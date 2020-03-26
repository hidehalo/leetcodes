package main

func defangIPaddr(address string) string {
	result := make([]rune, 0)
	for _, r := range address {
		if r == '.' {
			result = append(result, []rune("[.]")...)
			continue
		}
		result = append(result, r)
	}

	return string(result)
}
