package main

import "testing"

func Test_isToeplitzMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	"case1",
		// 	args{[][]int{
		// 		{1, 2, 3, 4},
		// 		{5, 1, 2, 3},
		// 		{9, 5, 1, 2},
		// 	}},
		// 	true,
		// },
		{
			"case2",
			args{[][]int{
				{1, 2},
				{2, 2},
			}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
