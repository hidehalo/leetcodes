package main

import (
	"testing"

	"../ds"
)

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ds.ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{ds.NewList([]int{1, 0, 1})},
			5,
		},
		{
			"case2",
			args{ds.NewList([]int{0})},
			0,
		},
		{
			"case3",
			args{ds.NewList([]int{0, 0})},
			0,
		},
		{
			"case4",
			args{ds.NewList([]int{1})},
			1,
		},
		{
			"case5",
			args{ds.NewList([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0})},
			18880,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
