package main

import "testing"

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{[]int{4, 2, 3}},
			true,
		},
		{
			"case2",
			args{[]int{4, 2, 1}},
			false,
		},
		{
			"case3",
			args{[]int{3, 4, 2, 3}},
			false,
		},
		{
			"case4",
			args{[]int{-1, 4, 2, 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
