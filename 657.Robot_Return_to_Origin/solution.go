package main

func judgeCircle(moves string) bool {
	h := 0
	v := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			h++
		} else if moves[i] == 'D' {
			h--
		} else if moves[i] == 'L' {
			v--
		} else if moves[i] == 'R' {
			v++
		}
	}
	return h == 0 && v == 0
}
