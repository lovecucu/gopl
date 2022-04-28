package ex23_test

import (
	ex23 "gopl/ch2/ex2.3-2.5"
	"testing"
)

func TestPopCount(t *testing.T) {
	// 全通过
	if ex23.PopCount(1) != 1 || ex23.PopCount(3) != 2 || ex23.PopCount(10) != 2 {
		t.Error("TestPopCount failed")
	}
	if ex23.PopCountCircle(1) != 1 || ex23.PopCountCircle(3) != 2 || ex23.PopCountCircle(10) != 2 {
		t.Error("TestPopCount failed")
	}
	if ex23.PopCountBit(1) != 1 || ex23.PopCountBit(3) != 2 || ex23.PopCountBit(10) != 2 {
		t.Error("TestPopCount failed")
	}
	if ex23.PopCountClearRight(1) != 1 || ex23.PopCountClearRight(3) != 2 || ex23.PopCountClearRight(10) != 2 {
		t.Error("TestPopCount failed")
	}
}

// 后续性能测试结果(go test -bench . -benchtime 30s)：
// popCount > popCountCircle > popCountClearRight > popCountBit
// BenchmarkPopCount-8                     1000000000               0.3123 ns/op
// BenchmarkPopCountCircle-8               1000000000               2.686 ns/op
// BenchmarkPopCountClearRight-8           1000000000               8.817 ns/op
// BenchmarkPopCountBit-8                  1000000000              25.79 ns/op
func benchmarkPopCount(b *testing.B, f ex23.PopCountFunc) {
	for i := 0; i < b.N; i++ {
		_ = f(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	benchmarkPopCount(b, ex23.PopCount)
}

func BenchmarkPopCountCircle(b *testing.B) {
	benchmarkPopCount(b, ex23.PopCountCircle)
}

func BenchmarkPopCountClearRight(b *testing.B) {
	benchmarkPopCount(b, ex23.PopCountClearRight)
}

func BenchmarkPopCountBit(b *testing.B) {
	benchmarkPopCount(b, ex23.PopCountBit)
}
