package main

import "testing"

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]int{1, 1, 2, 2, 3, 3}},
			3,
		},
		{
			"case2",
			args{[]int{1, 1, 2, 3}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
