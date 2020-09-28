package main

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{"abcdefg", 2},
			"bacdfeg",
		},
		{
			"case2",
			args{"a", 2},
			"a",
		},
		{
			"case3",
			args{"acacacacac", 4},
			"cacaacacca",
		},
		{
			"case4",
			args{"axxa", 3},
			"xxaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
