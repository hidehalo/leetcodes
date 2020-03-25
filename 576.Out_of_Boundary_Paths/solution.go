package main

import "fmt"

const (
	// LOMASK is bit mask of height
	LOMASK = 0x1f
	// HIMASK is bit mask of length
	HIMASK = 0x1f << 5
	// OPT code of move left
	LEFT = 0x0
	// OPT code of move right
	RIGHT = 0x1
	// OPT code of move up
	UP = 0x2
	// OPT code of move down
	DOWN = 0x3
)

func findPaths(m int, n int, N int, i int, j int) int {
	cp := make([]int, 0)
	// store boundary point which is validated
	for k := 0; k < m; k++ {
		if abs(i-k)+j+1 <= N {
			cp = append(cp, cood2bit(k, 0))
		}
		if abs(i-k)+n-j <= N {
			cp = append(cp, cood2bit(k, n-1))
		}
	}

	// duplication case 1 (m>1)
	// * 0 0 0 *
	// 0 0 0 0 0
	// * 0 0 0 *

	// duplication case 2 (m=1)
	// * * * * *

	for k := 0; k < n; k++ {
		// avoid duplication case 1
		if k == 0 || k == n-1 {
			continue
		}
		if i+abs(j-k)+1 <= N {
			cp = append(cp, cood2bit(0, k))
		}
		if m > 1 && // avoid duplication case 2
			m-i+abs(j-k) <= N {
			cp = append(cp, cood2bit(m-1, k))
		}
	}
	// dp := make(map[int][4]int)
	fmt.Printf("%v", cp)
	// todo design&impl&test
	// problem: is next move approve to boundary point? && has move reverse for next move?
	// how to memorized and what ???
	// T(N) = 4T(N-1) ??? sounds sad
	// left(N-1)
	// right(N-1)
	// up(N-1)
	// down(N-1)
	return 0
}

func cood2bit(i, j int) int {
	return i<<5 | j
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
