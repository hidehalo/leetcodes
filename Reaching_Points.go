package main

import (
	"fmt"
	// "math"
	// "strconv"
)

// var dp map[string]bool

// const THREADHOLD = 20

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for sx < tx && sy < ty {
		if tx < ty {
			ty %= tx
		} else {
			tx %= ty
		}
	}

	return sx == tx && (ty-sy)%sx == 0 || sy == ty && (tx-sx)%sy == 0
	// if sx == tx && sy == ty {
	// 	return true
	// }
	// a1 := strconv.Itoa(sx)
	// a2 := strconv.Itoa(sy)
	// b1 := strconv.Itoa(tx)
	// b2 := strconv.Itoa(ty)
	// sutiation := a1 + "," + a2 + "," + b1 + "," + b2
	// fmt.Println("situation:", sutiation)
	// if v, ok := dp[sutiation]; ok {
	// 	return v
	// }
	// r1 := []int{sx + sy, sy}
	// r2 := []int{sx, sx + sy}
	// var r1dead, r2dead bool
	// var d1, d2, d3, d4 float64
	// d1 = (float64)(sx - tx)
	// d2 = (float64)(r1[0] - tx)
	// d3 = (float64)(sy - ty)
	// d4 = (float64)(r2[1] - ty)
	// fmt.Println("distance:", math.Abs(d1) < math.Abs(d2), math.Abs(d3) < math.Abs(d4))
	// if math.Abs(d1) < math.Abs(d2) {
	// 	r1dead = true
	// }
	// if math.Abs(d3) < math.Abs(d4) {
	// 	r2dead = true
	// }
	// if r1dead && r2dead {
	// 	dp[sutiation] = false

	// 	return false
	// }
	// if r1dead != true && r1[0] == tx && r1[1] == ty {
	// 	dp[sutiation] = true

	// 	return true
	// } else if r2dead != true && r2[0] == tx && r2[1] == ty {
	// 	dp[sutiation] = true

	// 	return true
	// } else if r1dead == true && r2dead != true {
	// 	//speed up
	// 	if (ty-r2[1])/r2[0] > THREADHOLD {
	// 		r2[1] = ty - (ty-r2[1])%r2[0]
	// 	}
	// 	return reachingPoints(r2[0], r2[1], tx, ty)
	// } else if r1dead != true && r2dead == true {
	// 	//speed up
	// 	if (tx-r1[0])/r1[1] > THREADHOLD {
	// 		r1[0] = tx - (tx-r1[0])%r1[1]
	// 	}
	// 	return reachingPoints(r1[0], r1[1], tx, ty)
	// }
	// //speed up
	// if (ty-r2[1])/r2[0] > THREADHOLD {
	// 	r2[1] = (ty - (ty-r2[1])%r2[0]) / 1
	// 	if (tx-r2[0])/r2[1] > THREADHOLD {
	// 		r2[0] = (tx - (tx-r2[0])%r2[1]) / 1
	// 	}
	// }
	// //speed up
	// if (tx-r1[0])/r1[1] > THREADHOLD {
	// 	r1[0] = (tx - (tx-r1[0])%r1[1]) / 1
	// 	if (ty-r1[1])/r1[0] > THREADHOLD {
	// 		r1[1] = (ty - (ty-r1[1])%r1[0]) / 1
	// 	}
	// }

	// return reachingPoints(r1[0], r1[1], tx, ty) || reachingPoints(r2[0], r2[1], tx, ty)
}

func main() {
	// dp = make(map[string]bool)
	// fmt.Println(reachingPoints(35, 13, 455955547, 420098884))
	// fmt.Println(reachingPoints(30, 20, 300, 500))
	// fmt.Println(reachingPoints(44, 43, 921197891, 702724364))
	fmt.Println(reachingPoints(9, 10, 9, 10))
	// fmt.Println(dp)
}
