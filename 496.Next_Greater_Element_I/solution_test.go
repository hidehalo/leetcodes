package main

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			[]int{-1, 3, -1},
		},
		{
			"csae2",
			args{[]int{2, 4}, []int{1, 2, 3, 4}},
			[]int{3, -1},
		},
		{
			"case3",
			args{[]int{9, 8, 7}, []int{7, 8, 9, 10}},
			[]int{10, 9, 8},
		},
		{
			"case4",
			args{[]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}},
			[]int{7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
