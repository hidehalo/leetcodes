package main

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	"case1",
		// 	args{3},
		// 	0,
		// },
		// {
		// 	"case2",
		// 	args{5},
		// 	1,
		// },
		// {
		// 	"case3",
		// 	args{0},
		// 	0,
		// },
		{
			"case4",
			args{30},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
