package concat_test

import (
	"strings"
	"testing"
)

var args = []string{"hi", "there", "this", "is", "cucu", "13", "14"}

func concat(args []string) {
	r, sep := "", ""
	for _, a := range args {
		r += sep + a
		sep = " "
	}
}

func benchmarkConcat(n int) {
	for i := 0; i < n; i++ {
		concat(args)
	}
}

func benchmarkJoin(n int) {
	for i := 0; i < n; i++ {
		strings.Join(args, " ")
	}
}

func BenchmarkJoin10000(b *testing.B) {
	benchmarkJoin(10000)
}
func BenchmarkJoin1000000(b *testing.B) {
	benchmarkJoin(1000000)
}
func BenchmarkConcat10000(b *testing.B) {
	benchmarkConcat(10000)
}
func BenchmarkConcat1000000(b *testing.B) {
	benchmarkConcat(1000000)
}
