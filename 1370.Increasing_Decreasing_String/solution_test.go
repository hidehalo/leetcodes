package main

import "testing"

func Test_sortString(t *testing.T) {
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
			args{"aaaabbbbcccc"},
			"abccbaabccba",
		},
		{
			"case2",
			args{"rat"},
			"art",
		},
		{
			"case3",
			args{"leetcode"},
			"cdelotee",
		},
		{
			"case4",
			args{"ggggggg"},
			"ggggggg",
		},
		{
			"case5",
			args{"spo"},
			"ops",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
