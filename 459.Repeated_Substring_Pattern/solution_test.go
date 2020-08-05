package main

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{"abab"},
			true,
		},
		{
			"case2",
			args{"aba"},
			false,
		},
		{
			"case3",
			args{"abcabcabcabc"},
			true,
		},
		{
			"case4",
			args{"bb"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
