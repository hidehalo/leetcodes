package main

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{"hello"},
			"holle",
		},
		{
			"case2",
			args{"leetcode"},
			"leotcede",
		},
		{
			"case3",
			args{"aA"},
			"Aa",
		},
		{
			"case4",
			args{"race car"},
			"race car",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
