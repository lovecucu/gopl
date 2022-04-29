package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("12345"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	pre := n % 3
	if pre == 0 {
		pre = 3
	}
	var buf bytes.Buffer
	buf.WriteString(s[:pre])
	for i := pre; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
