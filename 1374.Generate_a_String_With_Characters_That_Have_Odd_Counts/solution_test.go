package main

import "testing"

func Test_generateTheString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{3},
			"aaa",
		},
		{
			"case2",
			args{4},
			"aaab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTheString(tt.args.n); got != tt.want {
				t.Errorf("generateTheString() = %v, want %v", got, tt.want)
			}
		})
	}
}
