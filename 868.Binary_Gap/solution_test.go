package main

import "testing"

func Test_binaryGap(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{22},
			2,
		},
		{
			"case2",
			args{5},
			2,
		},
		{
			"case3",
			args{6},
			1,
		},
		{
			"case4",
			args{8},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.N); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
