package main

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{[]int{1, 0, 0, 0, 1}, 1},
			true,
		},
		{
			"case2",
			args{[]int{1, 0, 0, 0, 1}, 2},
			false,
		},
		{
			"case3",
			args{[]int{0}, 1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
