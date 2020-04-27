package main

import "testing"

func Test_numSpecialEquivGroups(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{[]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}},
			3,
		},
		{
			"case2",
			args{[]string{"abc", "acb", "bac", "bca", "cab", "cba"}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSpecialEquivGroups(tt.args.A); got != tt.want {
				t.Errorf("numSpecialEquivGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
