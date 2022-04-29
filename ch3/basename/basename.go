package basename

import "strings"

func Basename(s string) string {
	// 保留最后一个/后的字符串
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func BasenameWithStrings(s string) string {
	slashIndex := strings.LastIndex(s, "/")
	s = s[slashIndex+1:]
	if dotIndex := strings.LastIndex(s, "."); dotIndex >= 0 {
		s = s[:dotIndex]
	}
	return s
}
