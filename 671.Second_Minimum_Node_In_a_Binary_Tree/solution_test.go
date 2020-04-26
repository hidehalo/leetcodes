package main

import (
	"testing"

	"../ds"
)

// func Test_minValue(t *testing.T) {
// 	type args struct {
// 		root *ds.TreeNode
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"case1",
// 			args{ds.NewBST([]int{7, 5, 3, 8, 1, 6, 2, 0})},
// 			0,
// 		},
// 		{
// 			"case2",
// 			args{ds.NewBST([]int{10, 1, 11, 2, 9, 4, 8, 6, 7, 12})},
// 			1,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := minValue(tt.args.root); got != tt.want {
// 				t.Errorf("minValue() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_min(t *testing.T) {
// 	type args struct {
// 		a int
// 		b int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			"case1",
// 			args{10, 100},
// 			10,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := min(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("min() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_findSecondMinimumValue(t *testing.T) {
	type args struct {
		root *ds.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{ds.NewBST([]int{10, 1, 11, 2, 9, 4, 8, 6, 7, 12})},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSecondMinimumValue(tt.args.root); got != tt.want {
				t.Errorf("findSecondMinimumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
