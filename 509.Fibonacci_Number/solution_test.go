package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{2},
			1,
		},
		{
			"case1",
			args{3},
			2,
		},
		{
			"case3",
			args{4},
			3,
		},
		{
			"case4",
			args{1},
			1,
		},
		{
			"case5",
			args{0},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.N); got != tt.want {
				t.Errorf("fib(%d) = %v, want %v", tt.args.N, got, tt.want)
			}
		})
	}
}
