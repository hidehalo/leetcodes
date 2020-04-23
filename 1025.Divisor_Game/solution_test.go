package main

import "testing"

func Test_divisorGame(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{2},
			true,
		},
		{
			"case2",
			args{3},
			false,
		},
		{
			"case3",
			args{1},
			false,
		},
		{
			"case4",
			args{4},
			true,
		},
		{
			"case5",
			args{5},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorGame(tt.args.N); got != tt.want {
				t.Errorf("divisorGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
