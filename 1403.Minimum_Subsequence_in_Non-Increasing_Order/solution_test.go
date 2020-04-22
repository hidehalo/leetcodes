package main

import (
	"reflect"
	"testing"
)

func Test_minSubsequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{4, 3, 10, 9, 8}},
			[]int{10, 9},
		},
		{
			"case2",
			args{[]int{4, 4, 7, 6, 7}},
			[]int{7, 7, 6},
		},
		{
			"case3",
			args{[]int{6}},
			[]int{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubsequence(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
