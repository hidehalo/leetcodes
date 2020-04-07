package main

func toLowerCase(str string) string {
	bytes := make([]byte, 0)
	d := byte('a' - 'A')
	for i := 0; i < len(str); i++ {
		if str[i] <= 'Z' && str[i] >= 'A' {
			bytes = append(bytes, str[i]+d)
		} else {
			bytes = append(bytes, str[i])
		}
	}

	return string(bytes)
}
