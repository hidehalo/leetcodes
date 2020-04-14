package main

import "testing"

func Test_findLucky(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]int{2, 2, 3, 4}},
			2,
		},
		{
			"case2",
			args{[]int{1, 2, 2, 3, 3, 3}},
			3,
		},
		{
			"case3",
			args{[]int{2, 2, 2, 3, 3}},
			-1,
		},
		{
			"case4",
			args{[]int{5}},
			-1,
		},
		{
			"case4",
			args{[]int{7, 7, 7, 7, 7, 7, 7}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLucky(tt.args.arr); got != tt.want {
				t.Errorf("findLucky() = %v, want %v", got, tt.want)
			}
		})
	}
}
