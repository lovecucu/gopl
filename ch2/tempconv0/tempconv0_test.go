package tempconv0_test

import (
	"fmt"
	"gopl/ch2/tempconv0"
	"testing"
)

func TestType(t *testing.T) {
	var c tempconv0.Celsius
	var f tempconv0.Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f == 0) // true
	// fmt.Println(f == c) // 编译错误：mismatched types
	fmt.Println(c == tempconv0.Celsius(f)) // true
}
