package main

import "testing"

func Test_findRadius(t *testing.T) {
	type args struct {
		houses  []int
		heaters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]int{1, 2, 3}, []int{2}},
			1,
		},
		{
			"case2",
			args{[]int{1, 2, 3, 4}, []int{1, 4}},
			1,
		},
		{
			"case3",
			args{[]int{1, 2, 3}, []int{100}},
			99,
		},
		{
			"case4",
			args{[]int{1, 2, 3}, []int{-100}},
			103,
		},
		{
			"case5",
			args{[]int{1, 5}, []int{2}},
			3,
		},
		{
			"case6",
			args{[]int{1}, []int{1, 2, 3, 4}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRadius(tt.args.houses, tt.args.heaters); got != tt.want {
				t.Errorf("findRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}
