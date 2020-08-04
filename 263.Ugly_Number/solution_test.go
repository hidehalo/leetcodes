package main

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{6},
			true,
		},
		{
			"case2",
			args{8},
			true,
		},
		{
			"case3",
			args{14},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.num); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
