package main

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	type testcase struct {
		name string
		arg  []int
		want []string
	}
	testcases := []testcase{
		{
			"case1",
			[]int{0, 1, 2, 4, 5, 7},
			[]string{"0->2", "4->5", "7"},
		},
		{
			"case2",
			[]int{0, 2, 3, 4, 6, 8, 9},
			[]string{"0", "2->4", "6", "8->9"},
		},
		{
			"case3",
			[]int{},
			[]string{},
		},
		{
			"case4",
			[]int{-1},
			[]string{"-1"},
		},
		{
			"case5",
			[]int{0},
			[]string{"0"},
		},
		{
			"case6",
			[]int{-2147483648, -2147483647, 2147483647},
			[]string{"-2147483648->-2147483647", "2147483647"},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := summaryRanges(testcase.arg)
			if reflect.DeepEqual(got, testcase.want) != true {
				t.Errorf("summaryRange() %s want %v, %v was got", testcase.name, testcase.want, got)
			}
		})
	}
}
