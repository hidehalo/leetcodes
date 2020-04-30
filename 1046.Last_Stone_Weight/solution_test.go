package main

import "testing"

func Test_lastStoneWeight(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	"case1",
		// 	args{[]int{2, 7, 4, 1, 8, 1}},
		// 	1,
		// },
		// {
		// 	"case2",
		// 	args{[]int{1, 2, 2}},
		// 	1,
		// },
		// {
		// 	"case3",
		// 	args{[]int{1, 3}},
		// 	2,
		// },
		// {
		// 	"case4",
		// 	args{[]int{1}},
		// 	1,
		// },
		// {
		// 	"case5",
		// 	args{[]int{}},
		// 	0,
		// },
		// {
		// 	"case6",
		// 	args{[]int{1, 1}},
		// 	0,
		// },
		{
			"case7",
			args{[]int{10, 5, 4, 10, 3, 1, 7, 8}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeight(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
