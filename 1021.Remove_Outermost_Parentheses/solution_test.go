package main

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
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
			args{"(()())(())"},
			"()()()",
		},
		{
			"case2",
			args{"(()())(())(()(()))"},
			"()()()()(())",
		},
		{
			"case3",
			args{"()()"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.S); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
