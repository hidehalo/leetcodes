package main

import (
	"reflect"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"case0 basic",
			args{
				[][]int{},
				[]int{2, 5},
			},
			[][]int{{2, 5}},
		},
		{
			"case1 left cross only",
			args{
				[][]int{{1, 3}, {6, 9}},
				[]int{2, 5},
			},
			[][]int{{1, 5}, {6, 9}},
		},
		{
			"case2 right cross only",
			args{
				[][]int{{1, 2}, {3, 4}},
				[]int{0, 1},
			},
			[][]int{{0, 2}, {3, 4}},
		},
		{
			"case3 equals or inside",
			args{
				[][]int{{1, 2}, {3, 4}},
				[]int{1, 2},
			},
			[][]int{{1, 2}, {3, 4}},
		},
		{
			"case4 contains",
			args{
				[][]int{{1, 2}, {6, 10}},
				[]int{5, 100},
			},
			[][]int{{1, 2}, {5, 100}},
		},
		{
			"case5 contains",
			args{
				[][]int{{1, 2}, {6, 10}},
				[]int{0, 100},
			},
			[][]int{{0, 100}},
		},
		{
			"case6 contains",
			args{
				[][]int{{1, 2}, {6, 10}},
				[]int{0, 5},
			},
			[][]int{{0, 5}, {6, 10}},
		},
		{
			"case7 never crossed",
			args{
				[][]int{{1, 2}, {9, 10}},
				[]int{4, 5},
			},
			[][]int{{1, 2}, {4, 5}, {9, 10}},
		},
		{
			"case8 merge and shrink crossed intervals",
			args{
				[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				[]int{4, 8},
			},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			"failed case 1",
			args{
				[][]int{{2, 5}, {6, 7}, {8, 9}},
				[]int{0, 1},
			},
			[][]int{{0, 1}, {2, 5}, {6, 7}, {8, 9}},
		},
		{
			"failed case 2",
			args{
				[][]int{{0, 7}, {8, 8}, {9, 11}},
				[]int{4, 13},
			},
			[][]int{{0, 13}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeIntervals(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"c1",
			args{
				[]int{1, 10},
				[]int{2, 3},
			},
			[]int{1, 10},
		},
		{
			"c2",
			args{
				[]int{1, 10},
				[]int{0, 1},
			},
			[]int{0, 10},
		},
		{
			"c3",
			args{
				[]int{1, 10},
				[]int{10, 13},
			},
			[]int{1, 13},
		},
		{
			"c4",
			args{
				[]int{1, 10},
				[]int{0, 13},
			},
			[]int{0, 13},
		},
		{
			"c5",
			args{
				[]int{1, 10},
				[]int{1, 10},
			},
			[]int{1, 10},
		},
		{
			"c6",
			args{
				[]int{1, 10},
				[]int{11, 12},
			},
			[]int{},
		},
		{
			"c7",
			args{
				[]int{1, 10},
				[]int{-10, -1},
			},
			[]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := mergeIntervals(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
