package main

import "testing"

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{[]int{1, 2, 2, 1, 1, 3}},
			true,
		},
		{
			"case2",
			args{[]int{1, 2}},
			false,
		},
		{
			"case3",
			args{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}