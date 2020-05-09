package main

import (
	"reflect"
	"testing"
)

func TestMyHashMap_Remove(t *testing.T) {
	type args struct {
		key int
	}
	caseOne := Constructor()

	caseTwo := Constructor()
	caseTwo.Put(1, 100)
	caseTwo.Put(1, 1)

	caseThree := Constructor()
	caseThree.Put(1025, 1)
	caseThree.Put(1, 1)

	tests := []struct {
		name string
		this *MyHashMap
		args args
	}{
		{
			"case1",
			&caseOne,
			args{1},
		},
		{
			"case2",
			&caseTwo,
			args{1},
		},
		{
			"case3",
			&caseThree,
			args{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.this
			this.Remove(tt.args.key)

			if this.Get(tt.args.key) != -1 {
				t.Errorf("MyHashMap.Remove() failed")
			}
		})
	}
}

func TestMyHashMap_Get(t *testing.T) {
	type args struct {
		key int
	}
	caseOne := Constructor()

	caseTwo := Constructor()
	caseTwo.Put(1, 100)

	caseThree := Constructor()
	caseThree.Put(1, 100)
	caseThree.Put(1, 1)

	tests := []struct {
		name string
		this *MyHashMap
		args args
		want int
	}{
		{
			"case1",
			&caseOne,
			args{1},
			-1,
		},
		{
			"case2",
			&caseTwo,
			args{1},
			100,
		},
		{
			"case3",
			&caseThree,
			args{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.this
			if got := this.Get(tt.args.key); got != tt.want {
				t.Errorf("MyHashMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyHashMap_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"case1",
			args{1, 100},
		},
		{
			"case2",
			args{10000, 100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			this.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); reflect.TypeOf(MyHashMap{}) != reflect.TypeOf(got) {
				t.Errorf("Constructor() = %v, want %v", reflect.TypeOf(got), reflect.TypeOf(MyHashMap{}))
			}
		})
	}
}
