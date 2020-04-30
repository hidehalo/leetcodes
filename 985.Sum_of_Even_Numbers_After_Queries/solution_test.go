package main

import (
	"reflect"
	"testing"
)

func Test_sumEvenAfterQueries(t *testing.T) {
	type args struct {
		A       []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{1, 2, 3, 4},
				[][]int{
					{1, 0},
					{-3, 1},
					{-4, 0},
					{2, 3},
				}},
			[]int{8, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumEvenAfterQueries(tt.args.A, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumEvenAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
