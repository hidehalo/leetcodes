package main

import (
	"reflect"
	"testing"
)

func Test_createTargetArray(t *testing.T) {
	type args struct {
		nums  []int
		index []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}},
			[]int{0, 4, 1, 3, 2},
		},
		{
			"case2",
			args{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}},
			[]int{0, 1, 2, 3, 4},
		},
		{
			"case3",
			args{[]int{4, 2, 4, 3, 2}, []int{0, 0, 1, 3, 1}},
			[]int{2, 2, 4, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTargetArray(tt.args.nums, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTargetArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
