package main

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case1",
			args{[]byte{'h', 'e', 'l', 'l', 'o'}},
		},
		{
			"case2",
			args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
		})
	}
}
