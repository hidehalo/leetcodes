package main

import "testing"

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		root *TreeNode
		L    int
		R    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{newBST([]int{10, 5, 15, 3, 7, 18}), 7, 15},
			32,
		},
		{
			"case2",
			args{newBST([]int{111, 61, 161, 35, 87, 137, 187, 23, 49, 75, 99, 125, 149, 175, 199, 17, 29, 43, 55, 69, 81, 93, 105, 119, 131, 143, 155, 169, 181, 193, 205, 13, 21, 27, 33, 39, 47, 53, 59, 65, 73, 79, 85, 91, 97, 103, 109, 115, 123, 129, 135, 141, 147, 153, 159, 165, 173, 179, 185, 191, 197, 203, 209, 11, 15, 19, 25, 31, 37, 41, 45, 51, 57, 63, 67, 71, 77, 83, 89, 95, 101, 107, 113, 117, 121, 127, 133, 139, 145, 151, 157, 163, 167, 171, 177, 183, 189, 195, 201, 207}), 149, 155},
			608,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.root, tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
