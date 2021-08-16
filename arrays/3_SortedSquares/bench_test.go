package test

import (
	"roman_study/arrays/3_SortedSquares/funcs"
	"testing"
)

func BenchmarkSample1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := []int{-4, -1, 0, 3, 10}
		_ = funcs.SortedSquares1(v)
	}
}
func BenchmarkSample2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := []int{-4, -1, 0, 3, 10}
		_ = funcs.SortedSquares2(v)
	}
}
