package popcount_test

import (
	"gopl/ch2/popcount"
	"testing"
)

func TestPopCount(t *testing.T) {
	if popcount.PopCount(1) != 1 || popcount.PopCount(3) != 2 || popcount.PopCount(10) != 2 {
		t.Error("TestPopCount failed")
	}
}
