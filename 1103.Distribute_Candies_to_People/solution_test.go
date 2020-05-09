package main

import (
	"reflect"
	"testing"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		candies    int
		num_people int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case1",
			args{7, 4},
			[]int{1, 2, 3, 1},
		},
		{
			"case2",
			args{10, 3},
			[]int{5, 2, 3},
		},
		{
			"case3",
			args{60, 4},
			[]int{15, 18, 15, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandies(tt.args.candies, tt.args.num_people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
