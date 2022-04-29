package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-12345.12"))
}

func comma(s string) string {
	var buf bytes.Buffer
	start := 0
	// 解决正负号问题
	if s[0] == '-' || s[0] == '+' {
		buf.WriteByte(s[0])
		start = 1
	}
	// 找出小数点位置，小数点前需要加逗号，小数位不需要
	end := strings.Index(s, ".")
	if end == -1 {
		end = len(s)
	}

	ms := s[start:end]

	// 去除小数和符号位后的计算
	pre := len(ms) % 3
	if pre == 0 {
		pre = 3
	}
	buf.Write([]byte(ms[:pre]))

	for i := pre; i < len(ms); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(ms[i : i+3])
	}

	// 小数位后
	buf.WriteString(s[end:])
	return buf.String()
}
