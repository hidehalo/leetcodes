package main

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		S string
		K int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{"5F3Z-2e-9-w", 4},
			"5F3Z-2E9W",
		},
		{
			"case1",
			args{"2-5g-3-J", 2},
			"2-5G-3J",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
