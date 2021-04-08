package main

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	testcases := []struct {
		name string
		arg  []int
		want []int
	}{
		{
			"case1",
			[]int{1, 2, 2, 4},
			[]int{2, 3},
		},
		{
			"case2",
			[]int{1, 1},
			[]int{1, 2},
		},
		{
			"case3",
			[]int{3, 2, 2},
			[]int{2, 1},
		},
		{
			"case4",
			[]int{2, 2},
			[]int{2, 1},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := findErrorNums(testcase.arg)
			if !reflect.DeepEqual(got, testcase.want) {
				t.Errorf("findErrorNums() %s got %v, but %v wanted", testcase.name, got, testcase.want)
			}
		})
	}
}
