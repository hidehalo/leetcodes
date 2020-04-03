package main

import "testing"

func BenchmarkIsEvenNumberOfDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEvenNumberOfDigits(1)
	}
}
