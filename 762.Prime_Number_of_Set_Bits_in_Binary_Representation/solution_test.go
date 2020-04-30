package main

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	type args struct {
		L int
		R int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{6, 10},
			4,
		},
		{
			"case2",
			args{10, 15},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimeSetBits(tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
