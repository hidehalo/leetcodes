package main

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{"A"},
			1,
		},
		{
			"case2",
			args{"AB"},
			28,
		},
		{
			"case3",
			args{"ZY"},
			701,
		},
		{
			"case4",
			args{"AZ"},
			52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
