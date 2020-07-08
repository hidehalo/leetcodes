package main

import "testing"

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]int{3, 2, 1}},
			1,
		},
		{
			"case2",
			args{[]int{1, 2}},
			2,
		},
		{
			"case3",
			args{[]int{2, 2, 3, 1}},
			1,
		},
		{
			"case4",
			args{[]int{1, 1, 2}},
			2,
		},
		{
			"case5",
			args{[]int{5, 2, 2}},
			5,
		},
		{
			"case6",
			args{[]int{1, 2, 2, 5, 3, 5}},
			2,
		},
		{
			"case7",
			args{[]int{1, 2, -2147483648}},
			-2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
