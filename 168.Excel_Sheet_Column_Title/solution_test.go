package main

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{1},
			"A",
		},
		{
			"case2",
			args{28},
			"AB",
		},
		{
			"case3",
			args{701},
			"ZY",
		},
		{
			"case4",
			args{52},
			"AZ",
		},
		{
			"case5",
			args{26},
			"Z",
		},
		{
			"case5",
			args{78},
			"BZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
