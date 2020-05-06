package main

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{[][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}},
			"Sao Paulo",
		},
		{
			"case2",
			args{[][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}},
			"A",
		},
		{
			"case3",
			args{[][]string{{"A", "Z"}}},
			"Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.args.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
