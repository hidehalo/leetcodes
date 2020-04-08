package main

import (
	"strconv"
)

var dp map[int]string

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if dp == nil {
		dp = make(map[int]string)
		dp[1] = "1"
	}
	if ret, ok := dp[n]; ok == true {
		return ret
	} else {
		for k := 2; k <= n; k++ {
			if ret, ok = dp[k]; ok == true {
				continue
			}
			ret = dp[k-1]
			size := len(ret)
			cur := ""
			for i := 0; i < size; i++ {
				count := 1
				for (i+1 < size) && (ret[i] == ret[i+1]) {
					count++
					i++
				}
				cur += strconv.Itoa(count) + string(ret[i])
			}
			dp[k] = cur
			ret = cur
		}
		return ret
	}
}
