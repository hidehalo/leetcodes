package main

func numUniqueEmails(emails []string) int {
	ret := 0
	uniq := make(map[string]int)
	for i := 0; i < len(emails); i++ {
		key := make([]byte, 0)
		cut := false
		ignore := false

		for j := 0; j < len(emails[i]); j++ {
			if emails[i][j] == '@' {
				cut = true
				ignore = false
			} else if ignore == true {
				continue
			} else if cut == false && emails[i][j] == '+' {
				ignore = true
				continue
			} else if cut == false && emails[i][j] == '.' {
				continue
			}
			key = append(key, emails[i][j])
		}

		keyStr := string(key)
		if uniq[keyStr] == 0 {
			uniq[keyStr] = 1
			ret++
		}
	}

	return ret
}
