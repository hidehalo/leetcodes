package main

import "testing"

func TestMaximumProduct(t *testing.T) {
	type testcase struct {
		name string
		arg  []int
		want int
	}
	testcases := []testcase{
		testcase{
			"case1",
			[]int{1, 2, 3},
			6,
		},
		testcase{
			"case2",
			[]int{1, 2, 3, 4},
			24,
		},
		testcase{
			"case3",
			[]int{-1, -2, -3},
			-6,
		},
		testcase{
			"case4",
			[]int{-1, -2, -3, 0, 1, 2},
			12,
		},
		testcase{
			"case5",
			[]int{-1, -2, -3, -4, -5, 0},
			0,
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			got := maximumProduct(testcase.arg)
			if got != testcase.want {
				t.Errorf("maximumProduct() %s got %d, but %d wanted", testcase.name, got, testcase.want)
			}
		})
	}
}
