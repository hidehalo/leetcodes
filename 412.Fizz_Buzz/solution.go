package main

import "strconv"

func fizzBuzz(n int) []string {
	ret := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ret = append(ret, "FizzBuzz")
		} else if i%3 == 0 {
			ret = append(ret, "Fizz")
		} else if i%5 == 0 {
			ret = append(ret, "Buzz")
		} else {
			ret = append(ret, strconv.FormatInt(int64(i), 10))
		}
	}

	return ret
}
