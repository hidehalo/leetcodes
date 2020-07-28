package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"case2",
			args{[]int{1, 2, 3, 4, 5, 6, 7}, 4},
			[]int{4, 5, 6, 7, 1, 2, 3},
		},
		{
			"case3",
			args{[]int{1, 2, 3, 4, 5, 6, 7}, 5},
			[]int{3, 4, 5, 6, 7, 1, 2},
		},
		{
			"case4",
			args{[]int{-1, -100, 3, 99}, 2},
			[]int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
