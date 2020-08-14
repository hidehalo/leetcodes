package main

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{"abc", "ahbgdc"},
			true,
		},
		{
			"case2",
			args{"axc", "ahbgdc"},
			false,
		},
		{
			"case3",
			args{"", "abv"},
			true,
		},
		{
			"case4",
			args{"acb", "ahbgdc"},
			false,
		},
		{
			"case5",
			args{"b", "c"},
			false,
		},
		{
			"case6",
			args{"aaaaaa", "bbaaaa"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
