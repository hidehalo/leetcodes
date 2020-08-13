package main

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{"1807", "7810"},
			"1A3B",
		},
		{
			"case2",
			args{"1123", "0111"},
			"1A1B",
		},
		{
			"case3",
			args{"00112233445566778899", "16872590340158679432"},
			"3A17B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
