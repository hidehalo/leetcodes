package main

import (
	"reflect"
	"testing"
)

func Test_prisonAfterNDays(t *testing.T) {
	type args struct {
		cells []int
		N     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{[]int{1, 0, 0, 1, 0, 0, 1, 0}, 1000000000},
			[]int{0, 0, 1, 1, 1, 1, 1, 0},
		},
		{
			"case2",
			args{[]int{0, 0, 1, 1, 1, 1, 0, 0}, 8},
			[]int{0, 0, 0, 1, 1, 0, 0, 0},
		},
		{
			"case3",
			args{[]int{0, 1, 0, 1, 1, 0, 0, 1}, 7},
			[]int{0, 0, 1, 1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prisonAfterNDays(tt.args.cells, tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prisonAfterNDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
