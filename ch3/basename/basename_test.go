package basename_test

import (
	"gopl/ch3/basename"
	"testing"
)

// 性能对比如下：使用strings包的效率较低
// BenchmarkBasename-8                     230201650                5.013 ns/op
// BenchmarkBasenameWithStrings-8          102867220               11.54 ns/op
func TestBaseName(t *testing.T) {
	if basename.Basename("a/b.c.go") != "b.c" || basename.Basename("a/b/c.go") != "c" {
		t.Error("TestBaseName failed")
	}

	if basename.BasenameWithStrings("a/b.c.go") != "b.c" || basename.BasenameWithStrings("a/b/c.go") != "c" {
		t.Error("TestBaseName failed")
	}
}

func BenchmarkBasename(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basename.Basename("a/b.c.go")
	}
}

func BenchmarkBasenameWithStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		basename.BasenameWithStrings("a/b.c.go")
	}
}
