package main

import (
	"fmt"
	"testing"

	"../ds"
)

func Test_increasingBST(t *testing.T) {
	type args struct {
		root *ds.TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case1",
			args{ds.NewBST([]int{10, 8, 11, 7, 9, 12, 2, 1, 20, 13, 5, 24})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := increasingBST(tt.args.root)
			p := got
			for p != nil {
				if p.Left != nil {
					fmt.Println("error")
				}
				fmt.Println(p.Val)
				p = p.Right
			}
		})
	}
}
