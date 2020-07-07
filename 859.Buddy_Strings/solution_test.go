package main

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	"case1",
		// 	args{"ab", "ba"},
		// 	true,
		// },
		// {
		// 	"case2",
		// 	args{"ab", "ab"},
		// 	false,
		// },
		{
			"case3",
			args{"aa", "aa"},
			true,
		},
		// {
		// 	"case4",
		// 	args{"aaaaaaabc", "aaaaaaacb"},
		// 	true,
		// },
		// {
		// 	"case5",
		// 	args{"", "aa"},
		// 	false,
		// },
		// {
		// 	"case6",
		// 	args{"", ""},
		// 	false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
