package main

import (
	"reflect"
	"testing"

	"../ds"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *ds.TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ds.TreeNode
	}{
		{
			"case1",
			args{ds.NewBST([]int{4, 2, 1, 3, 7}), 2},
			nil, // Can not test
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
