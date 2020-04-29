package main

import (
	"strconv"
)

func calPoints(ops []string) int {
	ac := make([]int, 0)
	ret := 0
	for i := 0; i < len(ops); i++ {
		if ops[i] == "+" {
			if len(ac) >= 2 {
				// fmt.Printf("2.1: (%d) + (%d) = %d\n", ac[len(ac)-2], ac[len(ac)-1], ac[len(ac)-2]+ac[len(ac)-1])
				ac = append(ac, ac[len(ac)-2]+ac[len(ac)-1])
				ret += ac[len(ac)-1]
			} else if len(ac) == 0 {
				// fmt.Printf("2.2: 0 + 0 = 0\n")
				ac = append(ac, 0)
			} else {
				// fmt.Printf("2.3: 0 + (%d) = %d\n", ac[0], ret+ac[0])
				ac = append(ac, ac[0])
				ret += ac[1]
			}
		} else if len(ac) > 0 && ops[i] == "C" {
			// fmt.Printf("4: %d - (%d) = %d\n", ret, ac[len(ac)-1], ret-ac[len(ac)-1])
			ret -= ac[len(ac)-1]
			ac = ac[:len(ac)-1]
		} else if ops[i] == "D" {
			if len(ac) > 0 {
				// fmt.Printf("3.1: %d * 2 = %d\n", ac[len(ac)-1], 2*ret)
				ac = append(ac, ac[len(ac)-1]*2)
				ret += ac[len(ac)-1]
			} else {
				// fmt.Printf("3.2: 0 * 2 = 0\n")
				ac = append(ac, 0)
			}
		} else {
			n, err := strconv.Atoi(ops[i])
			if err != nil {
				continue
			}
			// fmt.Printf("1: %d + (%d) = %d\n", ret, n, ret+n)
			ret += n
			ac = append(ac, n)
		}
	}
	return ret
}
