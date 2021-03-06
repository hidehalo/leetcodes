package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{"abbaca"},
			"ca",
		},
		{
			"case2",
			args{"azxxzy"},
			"ay",
		},
		{
			"case3",
			args{"aaaaaaaa"},
			"",
		},
		{
			"case4",
			args{"aababaab"},
			"ba",
		},
		{
			"case5",
			args{"bbaacaab"},
			"cb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.S); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
