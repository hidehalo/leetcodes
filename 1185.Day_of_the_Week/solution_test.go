package main

import "testing"

func Test_dayOfTheWeek(t *testing.T) {
	type args struct {
		day   int
		month int
		year  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{31, 8, 2019},
			"Saturday",
		},
		{
			"case2",
			args{18, 7, 1999},
			"Sunday",
		},
		{
			"case3",
			args{15, 8, 1993},
			"Sunday",
		},
		{
			"case4",
			args{1, 1, 1971},
			"Friday",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfTheWeek(tt.args.day, tt.args.month, tt.args.year); got != tt.want {
				t.Errorf("dayOfTheWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
