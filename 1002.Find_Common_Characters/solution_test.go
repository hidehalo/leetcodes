package main

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	type args struct {
		A []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"case1",
			args{[]string{"bella", "label", "roller"}},
			[]string{"l", "l", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonChars(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
