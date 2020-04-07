package main

import "testing"

func Test_balancedStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{"RLRRLLRLRL"},
			4,
		},
		{
			"case2",
			args{"RLLLLRRRLR"},
			3,
		},
		{
			"case3",
			args{"LLLLRRRR"},
			1,
		},
		{
			"case4",
			args{"RLRRRLLRLL"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
