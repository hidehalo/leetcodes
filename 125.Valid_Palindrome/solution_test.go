package main

import "testing"

func Test_isPalindrome(t *testing.T) {
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
			args{"A man, a plan, a canal: Panama"},
			true,
		},
		{
			"case2",
			args{"race a car"},
			false,
		},
		{
			"case3",
			args{"A"},
			true,
		},
		{
			"case4",
			args{""},
			true,
		},
		{
			"case5",
			args{" "},
			true,
		},
		{
			"case6",
			args{"0P"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
