package main

import "strconv"

func getHint(secret string, guess string) string {
	bulls := 0
	cowsMap := make([][]int, 2)
	for i := 0; i < 2; i++ {
		cowsMap[i] = make([]int, 10)
	}
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cowsMap[0][int(secret[i]-'0')]++
			cowsMap[1][int(guess[i]-'0')]++
		}
	}
	cows := 0
	for i := 0; i < 10; i++ {
		if cowsMap[0][i] > 0 && cowsMap[1][i] > 0 {
			if cowsMap[1][i] > cowsMap[0][i] {
				cows += cowsMap[0][i]
			} else {
				cows += cowsMap[1][i]
			}
		}
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
