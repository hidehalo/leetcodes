package main

import "testing"

func Test_minCostToMoveChips(t *testing.T) {
	type args struct {
		chips []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]int{1, 2, 3}},
			1,
		},
		{
			"case2",
			args{[]int{2, 2, 2, 3, 3}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostToMoveChips(tt.args.chips); got != tt.want {
				t.Errorf("minCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
