package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			"case2",
			args{[]int{1, 11, 111, 1111, 0, 0, 0, 0}, 4, []int{9, 10, 19, 20}, 4},
			[]int{1, 9, 10, 11, 19, 20, 111, 1111},
		},
		{
			"case3",
			args{[]int{1}, 1, []int{}, 0},
			[]int{1},
		},
		{
			"case4",
			args{[]int{}, 0, []int{}, 0},
			[]int{},
		},
		{
			"case5",
			args{[]int{}, 0, []int{1}, 1},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
