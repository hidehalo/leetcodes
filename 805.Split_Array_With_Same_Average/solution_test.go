package main

import (
	"testing"
)

func Test_splitArraySameAverage(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{[]int{5, 95, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
			true,
		},
		{
			"case2",
			args{[]int{2, 12, 18, 16, 19, 3}},
			false,
		},
		{
			"case3",
			args{[]int{3, 1, 2}},
			true,
		},
		{
			"case4",
			args{[]int{60, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30, 30}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArraySameAverage(tt.args.A); got != tt.want {
				t.Errorf("splitArraySameAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasKSum(t *testing.T) {
	type args struct {
		A      []int
		k      int
		target int
	}
	type testCase struct {
		name string
		args args
		want bool
	}
	tests := []testCase{
		{
			"[2,12,18,16,19,3]",
			args{[]int{2, 12, 18, 16, 19, 3}, 3, 35},
			false,
		},
		{
			"[3,1,2]",
			args{[]int{3, 1, 2}, 1, 2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasKSum(tt.args.A, tt.args.k, tt.args.target); got != tt.want {
				t.Errorf("hasKSum(%v) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
