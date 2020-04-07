package main

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{"Hello"},
			"hello",
		},
		{
			"case2",
			args{"here"},
			"here",
		},
		{
			"case3",
			args{"LOVELY"},
			"lovely",
		},
		{
			"case4",
			args{"al&phaBET"},
			"al&phabet",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.str); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
